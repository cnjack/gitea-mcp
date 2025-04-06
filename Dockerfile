FROM golang:1.24-alpine AS builder

ARG VERSION

WORKDIR /build

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=${VERSION}" -o gitea-mcp

FROM scratch

WORKDIR /app

COPY --from=builder /build/gitea-mcp .

CMD ["./gitea-mcp", "-t", "stdio"]