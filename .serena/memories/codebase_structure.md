# go-start Codebase Structure

## Top-Level Layout
```
go-start/
├── main.go                    # Entry point — calls cobra root command
├── go.mod / go.sum
├── cmd/
│   ├── root.go                # Cobra root command
│   └── new.go                 # `go-start new` subcommand (interactive wizard)
├── internal/
│   ├── ui/
│   │   └── console.go         # Console/UI helpers
│   └── generator/
│       ├── config.go          # Config struct (user selections)
│       ├── generator.go       # Core Generate() function + template rendering
│       ├── registry.go        # buildRegistry() — maps FileSpec entries (src template → dst path + condition)
│       └── gomod.go           # go.mod template generation
└── templates/                 # Go text/template files (.tmpl) for generated projects
    ├── core/                  # Always-included templates (models, pkg, config, etc.)
    ├── frameworks/
    │   ├── echo/              # Echo-specific delivery layer templates
    │   ├── gin/               # Gin-specific delivery layer templates
    │   └── fiber/             # Fiber-specific delivery layer templates
    ├── databases/             # DB connector templates (postgre, mysql, mongo, redis)
    ├── loggers/               # Logger implementation templates (zap, zerolog, logrus, slog)
    ├── auth/                  # JWT auth templates
    ├── swagger/               # Swagger/Scalar templates
    ├── tracing/               # OpenTelemetry templates
    └── migrations/            # Migration file templates (goose, golang-migrate)
```

## Key Types & Functions

### `internal/generator/config.go`
- `Config` struct — holds all user selections (ServiceName, ModuleName, Framework, HasPostgres, HasMySQL, HasMongoDB, HasRedis, UseGORM, MigrationTool, Logger, UseOTel, UsePrometheus, UseJWT, UseSwagger)
- Helper methods: `HasSQLDB()`, `HasAnyDB()`, `NeedsMigration()`, `UseGoose()`, `UseGolangMigrate()`

### `internal/generator/generator.go`
- `FileSpec` struct — `{SrcTemplate, DstPath, Condition func(*Config) bool}`
- `Generate(cfg *Config, templatesDir string) error` — iterates FileSpec registry, renders templates, runs `go mod tidy`
- `renderFile()`, `renderBytes()`, `renderString()` — template rendering helpers
- `templateFuncs()` — custom template functions (e.g., `ToLower`, `Title`, etc.)

### `internal/generator/registry.go`
- `buildRegistry(cfg *Config) []FileSpec` — returns all files to generate with conditions
- Helper predicates: `isEcho()`, `isGin()`, `isFiber()`

### `cmd/new.go`
- `runNew()` — drives the interactive wizard via `promptText`, `promptSelect`, `promptConfirm`
- `findTemplatesDir()` — locates the embedded/adjacent templates directory

## Templates
- All template files have `.tmpl` extension
- Templates use Go `text/template` syntax
- Template data is the `*Config` struct
- Custom template functions available from `templateFuncs()`
- Templates under `templates/core/` are always included
- Framework-specific templates are conditionally included based on `cfg.Framework`
