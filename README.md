# Hands-On RESTful Web Services with Go

This repository contains my implementations, modifications, and experiments based on the book **"Hands-On RESTful Web Services with Go"** by Naren Yellavula.

## 🚀 Projects & Chapters

Each directory represents a specific concept or chapter from the book, updated with custom Docker-based Swagger UI support.

- **[health-check](./health-check)**: Basic service to verify API availability and server time.
- **[mirrorFinder](./mirrorFinder)**: Service to find the fastest response times from a list of mirrors.
- **[multiple-handlers](./multiple-handlers)**: Demonstration of using `http.NewServeMux` with multiple endpoints.
- **[uuid-generator](./uuid-generator)**: A custom multiplexer implementation that generates hex-encoded UUIDs.

## 🛠 Development Workflow

### API Documentation
Each project includes an OpenAPI specification located in the `./api` directory. To view the interactive documentation:

1. Navigate to any project directory: `cd <project-name>`
2. Start the Swagger UI container: `./start_swagger_container.sh`
3. Open your browser to `http://localhost`

### Running the Services
To run any of the Go services locally:
```bash
go run main.go
# or
go run <filename>.go

