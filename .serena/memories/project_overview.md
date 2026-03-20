# go-start Project Overview

## Purpose
`go-start` is a **production-grade CLI scaffolding tool** for Golang clean architecture projects.
It generates fully-wired, production-ready Go REST APIs interactively via a CLI wizard.

## Tech Stack (the tool itself)
- **Language**: Go 1.23+
- **Module**: `go-start`
- **CLI framework**: `github.com/spf13/cobra`
- **Interactive prompts**: `github.com/manifoldco/promptui`
- **Spinner**: `github.com/briandowns/spinner`
- **Colors**: `github.com/fatih/color`

## What It Generates
Users pick from these options:
- **Framework**: Echo · Gin · Fiber
- **Databases** (multi-select): PostgreSQL · MySQL · MongoDB · Redis
- **ORM** (for SQL DBs): GORM · Raw SQL (pgx/sqlx)
- **Migration Tool**: goose · golang-migrate · none
- **Logger**: zap · zerolog · logrus · slog
- **Tracing**: OpenTelemetry · none
- **Metrics**: Prometheus · none
- **Auth**: JWT · none
- **Swagger**: yes · no

## Generated Project Architecture (Clean Architecture)
```
Delivery Layer (HTTP: Echo/Gin/Fiber)
  → Service Layer (business logic)
  → Repository Layer (data access: GORM/raw SQL/MongoDB/Redis)
  → Database(s)
```

## Project Repository
https://github.com/jeremygprawira/go-start
