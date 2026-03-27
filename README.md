# Hands-On RESTful Web Services with Go

This repository contains my implementations, modifications, and experiments based on the book **"Hands-On RESTful Web Services with Go"** by Naren Yellavula.

## 🚀 Projects & Chapters (Chapter 02)

All implementations for this chapter are located in the **[ch02/](./ch02)** directory, updated with custom Docker-based Swagger UI support.

- **[health-check](./ch02/health-check)**: Basic service to verify API availability and server time.
- **[lightweight-router](./ch02/lightweight-router)**: Implementation using `httprouter` to execute system commands (Go version, cat, ls).
- **[mirrorFinder](./ch02/mirrorFinder)**: Service to find the fastest response times from a list of mirrors.
- **[multiple-handlers](./ch02/multiple-handlers)**: Demonstration of using `http.NewServeMux` with multiple endpoints.
- **[mux-router](./ch02/mux-router)**: RESTful API using `gorilla/mux` with path-based (regex) and query-based matching.
- **[simple-static-file-server](./ch02/simple-static-file-server)**: Static file serving implementation using `httprouter`.
- **[uuid-generator](./ch02/uuid-generator)**: A custom multiplexer implementation that generates hex-encoded UUIDs.

## 🛠 Development Workflow

### API Documentation
Each project includes an OpenAPI specification located in the `./api` directory. To view the interactive documentation:

1. Navigate to a project directory: `cd ch02/<project-name>`
2. Start the Swagger UI container: `./start_swagger_container.sh`
3. Open your browser to `http://localhost`

### Running the Services
To run any of the Go services locally:
```bash
go run main.go
# or
go run <filename>.go

