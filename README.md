# gofiber-template

This is an opinionated template for building a RESTful API via HTTP transport project using [gofiber/fiber](https://github.com/gofiber/fiber) framework, with container support via [Docker](https://www.docker.com/) and CI/CD support via [GitHub Actions](https://github.com/features/actions).

## Features

- **Dependency Injection** via [go.uber.org/fx](https://github.com/uber-go/fx)
- **Configuration** via environment variable for easy container integration, utilizing
  - [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) for built-in common standard library value parsing support as well for custom value parsing support
  - [joho/godotenv](https://github.com/joho/godotenv) for using `.env` files to ease local development
- **Structured Logging** via [rs/zerolog](https://github.com/rs/zerolog) for zero-allocation JSON/[CBOR](https://github.com/rs/zerolog#binary-encoding) logging with support for log level, timestamp, and caller information, as well a human-readable render for local development
- **Fastest Web Framework** via [gofiber/fiber](https://github.com/gofiber/fiber) with zero memory allocation and performance in mind
- **Opinionated MVC Folder Structure** following [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- **Database Integration** via [uptrace/bun](https://github.com/uptrace/bun)
  - Built-in Database Migration via [uptrace/bun/migrate](https://bun.uptrace.dev/guide/migrations.html)
  - Easily write [complex queries](https://bun.uptrace.dev/#why-another-golang-orm), at the same time still have ready to use [struct-based CRUD operations](https://bun.uptrace.dev/guide/query-insert.html#example)
- **CI/CD** via [GitHub Actions](https://github.com/features/actions)
  - `.github/workflows/build-check.yml` for continuously checking the build status of the project by building the docker image of the project but not pushing it to the registry
  - `.github/workflows/build-release.yml` for releasing a version, in which GitHub Actions:
    1. Builds container image with automatically generated [OCI Image Format Specification labels](https://github.com/opencontainers/image-spec/blob/main/annotations.md) via [docker/metadata-action](https://github.com/docker/metadata-action)
    2. Pushes the image to [GitHub Container Registry](https://github.com/GalvinGao/gofiber-template/pkgs/container/gofiber-template)
    3. Creates a [GitHub Release](https://github.com/GalvinGao/gofiber-template/releases) with automatically generated [GitHub Release Notes](https://docs.github.com/en/repositories/releasing-projects-on-github/automatically-generated-release-notes) via [marvinpinto/action-automatic-releases](https://github.com/marvinpinto/action-automatic-releases)

## Getting Started

### 1. Get the template

#### `Use this template` button (recommended)

This repository is meant to be used as a template for your own project. You can use the `Use this template` button on the top right corner of this page to create your own repository from this template.

#### Clone this repository manually

If you want to clone this repository, you can do so by running the following command:

```bash
git clone git@github.com:GalvinGao/gofiber-template.git
```

### 2. Configure database URL in `.env`

```dotenv
DATABASE_URL=postgres://USERNAME:PASSWORD@localhost:5432/DATABASE_NAME?sslmode=disable
```

> More information about database URL can be found in [bun's documentation](https://bun.uptrace.dev/postgres/)

### 3. Initialize `bun` migration & Apply initial migrations

```bash
go run main.go db init
go run main.go db migrate
```

### 4. Launch

#### Install `gow` (optional)

[gow](https://github.com/mitranim/gow) stands for **Go W**atch. It is a tool that watches your Go source code and automatically recompiles and restarts your program when necessary. This allows you to see changes in real time which makes development much easier and faster.

```bash
go install github.com/mitranim/gow@latest
```

#### Start the server

After you install `gow` simply replace `go` with `gow` and start the server using:

```bash
gow run main.go start
```

## Debugging

### VSCode

Use the following `.vscode/launch.json` to launch the application with debugger attached:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Server",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/main.go",
      "args": ["start"]
    }
  ]
}
```

## License

[MIT License](LICENSE).

## Contributing

Issues and pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
