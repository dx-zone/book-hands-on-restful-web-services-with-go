#!/usr/bin/env bash

printf "Restarting Docker Swagger UI Container mapping: %s\n" "/app/api/openapi.json"
docker stop swaggerui 2>/dev/null || true

docker run --rm --name swaggerui -p 80:8080 \
  -e SWAGGER_JSON=/app/api/openapi.json \
  -v "$(pwd):/app" \
  swaggerapi/swagger-ui

