FROM golang:1.23.4-alpine AS base
WORKDIR /app

# builder
FROM base AS builder
ENV GOOS linux
ENV GOARCH amd64

# build-args
ARG VERSION

RUN apk --no-cache add bash git openssh

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
