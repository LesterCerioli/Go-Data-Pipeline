# go-pipeline

Go-Pipeline is a data pipeline API built with Fiber. Its packages follow the
layered structure used by FastBI-API-Go: HTTP handlers depend on service
contracts, implementations own persistence and business operations, and models
define the data model. FastBI-API-Go was inspected as a read-only architecture
reference; no files there were changed.

## Architecture

```text
cmd/                              application entry point
handlers/                          Fiber HTTP handlers
models/                            pipeline domain model
routers/                           route registration and API documentation endpoints
initializers/                     database, migrations, service dependency wiring
services/contracts/pipelines/     pipeline service interface
services/implementations/         application service implementations
cmd/docs/                          embedded Swagger JSON and YAML specification
```

The current API manages pipeline definitions. Source and destination settings
are JSON objects, leaving connector-specific configuration extensible.

## Run

Install Go 1.26 or newer, then run:

```sh
go mod tidy
go run ./cmd
```

The service listens on port `3000` by default. Set `PORT` to change it and
`DATABASE_DSN` to select a SQLite database file (defaults to `go-pipeline.db`).
The SQLite driver uses CGO, so a C compiler is required.

## API

- `GET /health`
- `GET /api/v1/pipelines`
- `POST /api/v1/pipelines`
- `GET /api/v1/pipelines/:id`
- `PUT /api/v1/pipelines/:id`
- `DELETE /api/v1/pipelines/:id`
- `GET /swagger.json` and `GET /swagger.yaml` — OpenAPI/Swagger 2.0 description

Create a pipeline with:

```json
{
  "name": "warehouse-ingest",
  "description": "Load events into the analytics warehouse",
  "source": {"type": "http", "url": "https://source.example/events"},
  "destination": {"type": "postgres", "table": "events"},
  "status": "draft"
}
```

Pipeline statuses are `draft`, `active`, and `paused`. Schema migrations are
registered in `initializers/migrations.go`.
