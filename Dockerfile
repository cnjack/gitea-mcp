# syntax=docker/dockerfile:1.4

# Build stage
FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

ARG VERSION=dev
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags="-s -w -X main.Version=${VERSION}" -o gitea-mcp

# Final stage
FROM gcr.io/distroless/static-debian11:nonroot

ENV GITEA_MODE=stdio

WORKDIR /app
COPY --from=builder --chown=nonroot:nonroot /app/gitea-mcp .

USER nonroot:nonroot

HEALTHCHECK --interval=30s --timeout=3s \
  CMD ["/app/gitea-mcp", "healthcheck"]

LABEL org.opencontainers.image.authors="your-team@example.com"
LABEL org.opencontainers.image.version="${VERSION}"

CMD ["/app/gitea-mcp"]
