# Build stage
FROM golang:1.26.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o testing-app .

# Run stage
FROM scratch

COPY --from=builder /app/testing-app /testing-app

EXPOSE 8080

ENTRYPOINT ["/testing-app"]
