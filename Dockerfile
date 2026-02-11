FROM golang:1.26.0-bookworm AS base
WORKDIR /app

# builder
FROM base AS builder

# build-args
ARG VERSION

RUN apt-get update && apt-get install -y --no-install-recommends bash git openssh-client && rm -rf /var/lib/apt/lists/*

# modules: utilize build cache
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

# inject versioning information & build the binary
RUN export BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ"); go build -o app -ldflags "-X github.com/GalvinGao/gofiber-template/internal/app/appbundle.Version=$VERSION -X github.com/GalvinGao/gofiber-template/internal/app/appbundle.BuildTimeString=$BUILD_TIME" .

# runner
FROM gcr.io/distroless/static-debian11 AS runner
WORKDIR /app

COPY --from=builder /app/app /app/app
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/app"]
