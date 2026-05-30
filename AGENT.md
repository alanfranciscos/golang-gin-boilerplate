# Engineering Principles & Architecture

This project follows strict engineering standards to ensure high quality, maintainability, and scalability.

## Principles

- **SOLID:** Every component is designed with a single responsibility, open for extension but closed for modification, and uses dependency inversion.
- **Clean Code:** Variable names are descriptive, functions are small, and the code is self-documenting.
- **Conventional Commits:** All commits follow the `type: description` format (e.g., `feat:`, `fix:`, `test:`, `chore:`).
- **Atomic Commits:** Each commit represents a single, logical change.

## Architecture: Hexagonal (Ports and Adapters)

The application is structured in layers to isolate the business logic from external concerns:

1.  **Domain Layer (`internal/domain`):** Contains the business models and entities. It has no dependencies on other layers.
2.  **Application Layer (`internal/application`):** Contains the use cases and port interfaces. It coordinates the flow of data to and from the domain.
3.  **Infrastructure Layer (`internal/infrastructure`):** Contains the adapters (e.g., Gin server, database drivers). It implements the interfaces defined in the application layer.

### Dependency Injection

Dependencies are injected via interfaces (Ports), allowing for easy testing with Mocks and decoupling from specific implementations.

## Quality Standards

- **Unit Testing:** Mandatory for all files.
- **Coverage:** Minimum of 90% coverage per file.
- **Static Analysis:** Code must pass linting and formatting checks.
