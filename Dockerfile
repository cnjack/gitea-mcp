# Build stage
FROM golang:1.24-bullseye AS builder

ARG VERSION

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${VERSION}" -o gitea-mcp

# Final stage
FROM debian:bullseye-slim

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apt-get update && \
    apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Create a non-root user
RUN useradd -r -u 1000 -m gitea-mcp

COPY --from=builder --chown=1000:1000 /app/gitea-mcp .

# Use the non-root user
USER gitea-mcp

CMD ["/app/gitea-mcp", "-t", "stdio"]