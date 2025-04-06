FROM golang:1.24-alpine AS builder

ARG VERSION

WORKDIR /build

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=${VERSION}" -o gitea-mcp

FROM ubuntu:24.04

WORKDIR /app

RUN apt-get update \
    && apt-get install ca-certificates --no-install-recommends -y

COPY --from=builder /build/gitea-mcp .

CMD ["./gitea-mcp", "-t", "stdio"]