# Project Architecture

This document describes the architectural patterns and components used in the `go-boilerplate` project.

## Hexagonal Architecture (Ports and Adapters)

The project follows the Hexagonal Architecture pattern to decouple the business logic from infrastructure and external services.

```mermaid
graph TD
    subgraph Primary_Adapters [Primary Adapters - Input]
        GIN[Gin Web Server]
        CLI[Main CLI/Entrypoint]
    end

    subgraph Application_Layer [Application Layer - Use Cases]
        HS[Health Check Service]
    end

    subgraph Domain_Layer [Domain Layer - Entities]
        HD[Health Model]
    end

    subgraph Secondary_Adapters [Secondary Adapters - Output]
        OTEL[OpenTelemetry Exporter]
    end

    CLI --> HS
    GIN --> HS
    HS --> HD
    HS --> OTEL
```

## Data Flow & Observability

```mermaid
sequenceDiagram
    participant User
    participant Gin as Gin Middleware (OTel)
    participant UC as Use Case
    participant SDK as OTel SDK

    User->>Gin: GET /health
    activate Gin
    Gin->>SDK: Start Span
    Gin->>UC: GetHealth(ctx)
    activate UC
    UC->>SDK: Add Event/Attribute
    UC-->>Gin: Health Data
    deactivate UC
    Gin->>SDK: End Span
    Gin-->>User: 200 OK (JSON)
    deactivate Gin
```

## Package Structure

- `cmd/api/`: Entry point of the application.
- `internal/application/`: Business logic and port interfaces.
- `internal/domain/`: Enterprise models and entities.
- `internal/infrastructure/`: Concrete implementations (Adapters).
    - `web/`: Gin server and handlers.
    - `telemetry/`: OpenTelemetry setup.
    - `config/`: Configuration management.
- `pkg/`: Shared libraries that are safe to use by other projects.
