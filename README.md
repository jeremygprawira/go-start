<div align="center">

<h1>🚀 go-start</h1>

<p><strong>A production-grade CLI scaffolding tool for Golang clean architecture projects</strong></p>

<p>
  <a href="https://golang.org/doc/go1.23"><img src="https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat-square&logo=go" alt="Go Version"/></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue?style=flat-square" alt="License"/></a>
  <img src="https://img.shields.io/badge/status-active-success?style=flat-square" alt="Status"/>
</p>

<p>
  Scaffold a complete, production-ready Go REST API in seconds — pick your framework, databases, logger, auth, tracing, and migrations interactively. Everything is wired and ready to code.
</p>

</div>

---

## ✨ Features

- 🧙 **Interactive wizard** — step-by-step prompts with a color-coded summary before generation
- ⚡ **Clean Architecture** — mirrors [go-echo-boilerplate](https://github.com/jeremygprawira/go-echo-boilerplate): delivery → service → repository layers
- 🔧 **Multi-framework** — Echo, Gin, or Fiber (all fully wired, not just stubs)
- 🗄️ **Multi-database** — PostgreSQL, MySQL, MongoDB, Redis (mix and match)
- 🔄 **ORM flexibility** — GORM or Raw SQL (pgx / sqlx) per SQL database
- 📦 **4 loggers** — Zap, Zerolog, Logrus, or slog (stdlib), all behind a unified interface
- 🔬 **Distributed tracing** — OpenTelemetry (OTLP/gRPC) optional
- 🔐 **JWT auth** — access + refresh tokens, middleware wired per framework
- 📖 **Swagger** — swaggo annotations + route registered (per framework)
- 🐳 **Docker ready** — Dockerfile, docker-compose with only your chosen database services
- 🔃 **Migrations** — Goose or golang-migrate SQL files scaffolded automatically
- 📝 **Dynamic go.mod** — only imports libraries matching your selections

---

## 📋 Option Matrix

| Choice | Options |
|---|---|
| **Framework** | `echo` · `gin` · `fiber` |
| **Database** | `PostgreSQL` · `MySQL` · `MongoDB` · `Redis` *(multi-select)* |
| **ORM** *(SQL DBs)* | `GORM` · `Raw SQL (pgx/sqlx)` |
| **Migration Tool** *(SQL DBs)* | `goose` · `golang-migrate` · `none` |
| **Logger** | `zap` · `zerolog` · `logrus` · `slog` |
| **Tracing** | `OpenTelemetry` · `none` |
| **Auth** | `JWT` · `none` |
| **Swagger** | `yes` · `no` |

---

## 📥 Installation

### Option A — Build from source

```bash
git clone https://github.com/jeremygprawira/go-start.git
cd go-start
go build -o go-start .
sudo mv go-start /usr/local/bin/   # optional: make globally available
```

### Option B — Run directly with `go run`

```bash
git clone https://github.com/jeremygprawira/go-start.git
cd go-start
go run main.go new
```

**Requirements:**
- Go 1.23+
- Docker & Docker Compose (to run generated projects)

---

## 🚀 Quick Start

```bash
go-start new
```

You'll walk through a series of prompts:

```
  ╔═══════════════════════════════════════╗
  ║        🚀  go-start  CLI              ║
  ║    Golang Clean Architecture Scaffold  ║
  ╚═══════════════════════════════════════╝

? Service Name: my-api
? Go Module Name: github.com/yourname/my-api
? HTTP Framework: echo
? Select Databases (e.g. 1,3):
  [1] PostgreSQL
  [2] MySQL
  [3] MongoDB
  [4] Redis
  Select: 1,4
? Database Access Layer: GORM (ORM)
? Migration Tool: goose
? Logger: zap
? Distributed Tracing: none
? Authentication: JWT
? API Documentation: Swagger (swaggo)

  ┌─ Project Summary ────────────────────────┐
  │  Service Name:   my-api                   │
  │  Go Module:      github.com/yourname/my-api│
  │  Framework:      ECHO                     │
  │  Databases:      PostgreSQL + Redis        │
  │  ORM:            GORM                     │
  │  Migrations:     goose                    │
  │  Logger:         zap                      │
  │  Tracing:        none                     │
  │  Auth:           JWT                      │
  │  Swagger:        yes (swaggo)             │
  └───────────────────────────────────────────┘

? Generate project? y

✅ Project generated successfully!
```

Then:

```bash
cd my-api
cp config/config.local.example.yaml config/config.local.yaml
# Edit config.local.yaml with your DB credentials
docker compose up -d
make run
```

---

## 📁 Generated Project Structure

```
my-api/
├── builds/
│   └── Dockerfile                  # multi-stage build
├── cmd/
│   └── http/
│       └── main.go                 # graceful lifecycle management
├── config/
│   └── config.local.example.yaml  # template config (copy and fill)
├── docs/                           # swagger docs (if enabled)
├── internal/
│   ├── config/
│   │   ├── config.go               # Viper-based loader (ENV + YAML)
│   │   └── model.go                # Configuration struct
│   ├── core/
│   │   ├── setup.go                # wires logger → db → repo → service → handler
│   │   └── teardown.go             # graceful shutdown cleanup
│   ├── deliveries/
│   │   └── http/
│   │       ├── router.go           # route registration + swagger + health
│   │       ├── api/v1/
│   │       │   ├── v1_handler.go   # API group registrar
│   │       │   └── user_v1_handler.go  # CRUD handler scaffold
│   │       ├── health_check/
│   │       │   └── health_handler.go
│   │       └── middleware/
│   │           ├── middleware.go   # middleware registrar
│   │           ├── cors.go
│   │           ├── logger.go       # request logging via app logger
│   │           ├── recover.go
│   │           ├── api_key.go      # X-API-Key header check
│   │           └── jwt.go          # JWT Bearer validation (if enabled)
│   ├── models/
│   │   ├── common_model.go         # BaseModel with UUID PK
│   │   ├── response_model.go       # standard API envelope
│   │   ├── health_model.go
│   │   └── user_model.go           # User entity + request/response DTOs
│   ├── pkg/
│   │   ├── database/
│   │   │   ├── database.go         # Database struct + Connect()
│   │   │   ├── postgres.go         # PostgreSQL connector (if selected)
│   │   │   ├── mysql.go            # MySQL connector (if selected)
│   │   │   ├── mongo.go            # MongoDB connector (if selected)
│   │   │   └── redis.go            # Redis connector (if selected)
│   │   ├── graceful/               # OS signal + shutdown orchestration
│   │   ├── logger/                 # chosen logger (unified Logger interface)
│   │   ├── jwtc/                   # JWT access/refresh token gen + parsing
│   │   ├── tracer/                 # OpenTelemetry tracer (if enabled)
│   │   ├── validator/              # go-playground/validator wrapper
│   │   ├── errorc/                 # structured AppError type
│   │   ├── response/               # framework-specific response helpers
│   │   ├── formatter/              # string/email formatters
│   │   └── generator/              # UUID + token generators
│   ├── repository/
│   │   ├── main_repository.go      # Repository aggregator
│   │   ├── pgsql/                  # PostgreSQL repos (if selected)
│   │   ├── mysql/                  # MySQL repos (if selected)
│   │   ├── mongo/                  # MongoDB repos (if selected)
│   │   └── cache/                  # Redis cache repo (if selected)
│   └── service/
│       ├── main_service.go         # Service aggregator + Dependencies
│       ├── health_service.go
│       └── user_service.go         # CRUD business logic scaffold
├── migration/
│   └── db/
│       ├── postgre/001_init.sql    # users table (goose/migrate format)
│       └── mysql/001_init.sql
├── .air.toml                       # hot-reload config
├── .env.example
├── .gitignore
├── docker-compose.yml              # only includes your selected DB services
├── Makefile
└── README.md
```

---

## 🔌 Design Patterns

### Clean Architecture Layers

```
Client Request
      │
      ▼
┌─────────────────────────┐
│   Delivery Layer        │  router.go + middleware + handler
│   (HTTP Framework)      │  ← framework-specific (Echo/Gin/Fiber)
└───────────┬─────────────┘
            │  service.Method(ctx, dto)
            ▼
┌─────────────────────────┐
│   Service Layer         │  business logic, validation, error handling
│   (Use Cases)           │
└───────────┬─────────────┘
            │  repository.Action(ctx, entity)
            ▼
┌─────────────────────────┐
│   Repository Layer      │  data access (GORM / raw SQL / MongoDB / Redis)
│   (Data Access)         │
└───────────┬─────────────┘
            │
            ▼
       Database(s)
```

### Unified Logger Interface

All 4 logger variants implement the same interface — you can swap loggers without touching any application code:

```go
type Logger interface {
    Info(ctx context.Context, msg string, fields ...Field)
    Error(ctx context.Context, msg string, fields ...Field)
    Warn(ctx context.Context, msg string, fields ...Field)
    Debug(ctx context.Context, msg string, fields ...Field)
    Fatal(ctx context.Context, msg string, fields ...Field)
}
```

### Wiring Pattern (core/setup.go)

```go
logger.Initialize(cfg)                     // 1. logger first
db, _ := database.Connect(cfg)             // 2. databases
repo := repository.New(db)                 // 3. repositories
svc  := service.New(service.Dependencies{ // 4. services
    Repository: *repo,
    JWTConfig:  jwtConfig,
    Config:     cfg,
})
handler.New(e, svc, cfg, jwtConfig)        // 5. HTTP layer last
```

---

## 🛠 Development Commands (Generated Project)

```bash
make run          # go run ./cmd/http/main.go
make build        # build binary to ./builds/bin/
make test         # go test ./... -v -cover
make lint         # golangci-lint run ./...
make tidy         # go mod tidy
make swagger      # swag init (regenerate docs)
make migrate-up   # run pending migrations (goose/golang-migrate)
make migrate-down # rollback last migration
make docker-up    # docker compose up --build -d
make docker-down  # docker compose down
make docker-logs  # tail app container logs
```

---

## 🐳 Docker

The generated `docker-compose.yml` only includes services matching your selected databases:

```bash
docker compose up -d          # start all services
docker compose down           # stop all services
docker compose logs -f app    # stream app logs
```

The `builds/Dockerfile` uses a two-stage build:
1. **Builder** — compiles the binary in a full Go image  
2. **Runner** — copies binary into a minimal Alpine image (`~10MB` final image)

---

## 📖 API Endpoints (Scaffold)

After generation, the following routes are registered out of the box:

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| `GET` | `/` | — | Root ping |
| `GET` | `/health` | — | Health check |
| `GET` | `/swagger/*` | — | Swagger UI *(if enabled)* |
| `POST` | `/api/v1/users` | API Key | Create user |
| `GET` | `/api/v1/users/:id` | API Key + JWT | Get user by ID |
| `PUT` | `/api/v1/users/:id` | API Key + JWT | Update user |
| `DELETE` | `/api/v1/users/:id` | API Key + JWT | Delete user |

> Headers: `X-API-Key: <key>` for API key auth · `Authorization: Bearer <token>` for JWT

---

## 🧩 Extending the Scaffold

### Add a new domain (e.g. `product`)

1. **Model** → `internal/models/product_model.go`
2. **Repository interface** → `internal/repository/pgsql/product_pgsql_repository.go`
3. **Register in main repo** → `internal/repository/main_repository.go`
4. **Service** → `internal/service/product_service.go`
5. **Register in main service** → `internal/service/main_service.go`
6. **Handler** → `internal/deliveries/http/api/v1/product_v1_handler.go`
7. **Register route** → `internal/deliveries/http/api/v1/v1_handler.go`
8. **Migration** → `migration/db/postgre/002_products.sql`

---

## 📦 Technology References

| Library | Purpose |
|---|---|
| [github.com/labstack/echo/v4](https://echo.labstack.com/) | Echo web framework |
| [github.com/gin-gonic/gin](https://gin-gonic.com/) | Gin web framework |
| [github.com/gofiber/fiber/v2](https://gofiber.io/) | Fiber web framework |
| [gorm.io/gorm](https://gorm.io/) | GORM ORM |
| [github.com/jackc/pgx/v5](https://github.com/jackc/pgx) | PostgreSQL raw driver |
| [github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx) | MySQL raw driver |
| [go.mongodb.org/mongo-driver/v2](https://www.mongodb.com/docs/drivers/go/) | MongoDB driver |
| [github.com/redis/go-redis/v9](https://github.com/redis/go-redis) | Redis client |
| [github.com/pressly/goose/v3](https://github.com/pressly/goose) | Goose migrations |
| [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate) | golang-migrate |
| [go.uber.org/zap](https://github.com/uber-go/zap) | Zap logger |
| [github.com/rs/zerolog](https://github.com/rs/zerolog) | Zerolog logger |
| [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) | Logrus logger |
| [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt) | JWT auth |
| [go.opentelemetry.io/otel](https://opentelemetry.io/) | OpenTelemetry tracing |
| [github.com/swaggo/swag](https://github.com/swaggo/swag) | Swagger docs |
| [github.com/spf13/viper](https://github.com/spf13/viper) | Config loader |
| [github.com/spf13/cobra](https://github.com/spf13/cobra) | CLI framework |

---

## 📄 License

MIT © 2025 [Jeremy Gerald Prawira](https://github.com/jeremygprawira)