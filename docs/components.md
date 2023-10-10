Here's a translation of the information about the components used in NeoShowcase:

## Standalone Components

### Buildkit
- Description: A tool for building Docker images. It can build Docker images without Docker and run them within containers. NeoShowcase uses it as a server for application builds, controlled remotely through gRPC API by starting a daemon called `buildkitd`.
- Note: It internally uses LLB (Low Level Build definition format) for image building, although Dockerfile is more commonly used.

### Buildpacks
- Description: A tool to build Docker images from application code without Dockerfiles. It detects optimal build settings from the code and creates Docker images. NeoShowcase mainly uses the Paketo Buildpacks implementation.
  
### gRPC
- Description: A Google-developed RPC framework using HTTP/2. NeoShowcase uses gRPC for communication between various purpose-specific servers, such as management and build servers.
  
### Connect
- Description: A protocol for communication using HTTP/1.1 and HTTP/2 POST methods. It is used for routing authentication to the traefik forward auth middleware in NeoShowcase.

### Traefik Proxy
- Description: A modern reverse proxy used for connecting components and routing user applications in NeoShowcase, especially in Kubernetes.

### protoc (Protocol Buffer Compiler)
- Description: A compiler used to generate `.go` files from `.proto` files for Protocol Buffers. It's used for defining the specifications of functions for communication between various components in NeoShowcase.

### evans
- Description: A gRPC client for interactive debugging and testing of gRPC APIs. It's used when direct access to gRPC is needed.

### sqldef
- Description: A tool for idempotent schema management of MySQL, PostgreSQL, SQLite3, and SQL Server using SQL files. It's used for managing database schema changes in a safe and repeatable manner.

### sql-migrate (currently not used, replaced by sqldef)
- Description: A database schema migration tool used to version control and apply database changes during development. NeoShowcase transitioned from sql-migrate to sqldef for better idempotence.

### tbls
- Description: A tool for generating documentation and linting from real databases. NeoShowcase uses it to automatically generate ER diagrams and documents from its database.

### golangci-lint
- Description: A linter for Go code, used for code formatting and detecting potential issues.

### Swagger / OpenAPI 3.0 (currently not used, replaced by Connect)
- Description: A specification for describing HTTP APIs. NeoShowcase used it for defining API specifications in the past but currently relies on `.proto` files for all communication.

### spectral (currently not used)
- Description: A linter for `swagger.yaml` files.

## Go Libraries

### sqlboiler
- Description: A code generator for SQL database ORMs in Go. NeoShowcase uses it to generate ORM libraries based on the database schema.

### Echo
- Description: A web server library used in NeoShowcase for various purposes, including serving static sites.

### logrus
- Description: A library for logging console output in a structured manner, used for NeoShowcase's logging.

### cobra / viper
- Description: Cobra is a command-line tool library, and Viper is a library for handling configuration files. They are used for creating command-line tools in NeoShowcase.

### Wire
- Description: A library for generating dependency injection code automatically. NeoShowcase uses it to handle dependency injection in its codebase.

### docker/client
- Description: The official Go library for interacting with Docker. NeoShowcase uses it to interact with Docker for container-related operations.

### Hub (currently not used)
- Description: A pub/sub-style internal event bus library used for communication between different parts of the codebase. It helps decouple components in the code but should be used cautiously to avoid overcomplicating control flow.

This translation should help you understand the various components and libraries used in NeoShowcase. If you have any specific questions or need more details about any component, feel free to ask.
