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

## Guardrails

To maintain the integrity of the project, the following rules are strictly enforced:

1.  **Architectural Isolation:** The Domain layer MUST NOT import any external frameworks (e.g., Gin, Gorm) or infrastructure-specific packages.
2.  **Commit Standards:**
    - Must follow **Conventional Commits**.
    - Must be in **English**.
    - Maximum of **70 characters** for the subject line.
    - **No co-authors** (`Co-authored-by:`) allowed.
    - Commits must be **atomic**.
3.  **Quality Enforcement:**
    - Every `.go` file must have a corresponding `_test.go` file.
    - `make test` will fail if coverage drops below 90% on any modified file.
    - Multi-stage Docker builds must be used to keep production images lean.
    - **Linting:** Code MUST pass `golangci-lint` without any errors or warnings.
    - **Formatting:** Code MUST be formatted using `go fmt`.


4.  **Database Access:**
    - **No ORM:** The use of ORMs (e.g., Gorm, Ent) is strictly prohibited.
    - **Raw SQL:** All database queries MUST be written as raw SQL to ensure maximum performance, predictability, and control.


