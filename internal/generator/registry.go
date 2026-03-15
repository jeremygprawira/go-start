package generator

// buildRegistry returns all FileSpecs for the project, with conditions.
// Templates are relative to the templates/ directory.
func buildRegistry(cfg *Config) []FileSpec {
	sn := cfg.ServiceName // shorthand
	_ = sn

	return []FileSpec{
		// ─── Core / Always Generated ────────────────────────────────────────
		{SrcTemplate: "core/cmd/http/main.go.tmpl", DstPath: "[[.ServiceName]]/cmd/http/main.go"},
		{SrcTemplate: "core/internal/config/config.go.tmpl", DstPath: "[[.ServiceName]]/internal/config/config.go"},
		{SrcTemplate: "core/internal/config/model.go.tmpl", DstPath: "[[.ServiceName]]/internal/config/model.go"},
		{SrcTemplate: "core/internal/core/setup.go.tmpl", DstPath: "[[.ServiceName]]/internal/core/setup.go"},
		{SrcTemplate: "core/internal/core/teardown.go.tmpl", DstPath: "[[.ServiceName]]/internal/core/teardown.go"},
		{SrcTemplate: "core/internal/models/common_model.go.tmpl", DstPath: "[[.ServiceName]]/internal/models/common_model.go"},
		{SrcTemplate: "core/internal/models/response_model.go.tmpl", DstPath: "[[.ServiceName]]/internal/models/response_model.go"},
		{SrcTemplate: "core/internal/models/health_model.go.tmpl", DstPath: "[[.ServiceName]]/internal/models/health_model.go"},
		{SrcTemplate: "core/internal/models/user_model.go.tmpl", DstPath: "[[.ServiceName]]/internal/models/user_model.go"},
		{SrcTemplate: "core/internal/pkg/graceful/graceful.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/graceful.go"},
		{SrcTemplate: "core/internal/pkg/validator/validator.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/validator.go"},
		{SrcTemplate: "core/internal/pkg/errorc/errorc.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/errorc/errorc.go"},
		{SrcTemplate: "core/internal/pkg/response/response.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/response/response.go"},
		{SrcTemplate: "core/internal/pkg/formatter/formatter.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/formatter/formatter.go"},
		{SrcTemplate: "core/internal/pkg/generator/generator.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/generator.go"},
		{SrcTemplate: "core/internal/repository/main_repository.go.tmpl", DstPath: "[[.ServiceName]]/internal/repository/main_repository.go"},
		{SrcTemplate: "core/internal/service/main_service.go.tmpl", DstPath: "[[.ServiceName]]/internal/service/main_service.go"},
		{SrcTemplate: "core/internal/service/health_service.go.tmpl", DstPath: "[[.ServiceName]]/internal/service/health_service.go"},
		{SrcTemplate: "core/internal/service/user_service.go.tmpl", DstPath: "[[.ServiceName]]/internal/service/user_service.go"},
		{SrcTemplate: "core/Makefile.tmpl", DstPath: "[[.ServiceName]]/Makefile"},
		{SrcTemplate: "core/gitignore.tmpl", DstPath: "[[.ServiceName]]/.gitignore"},
		{SrcTemplate: "core/air.toml.tmpl", DstPath: "[[.ServiceName]]/.air.toml"},
		{SrcTemplate: "core/README.md.tmpl", DstPath: "[[.ServiceName]]/README.md"},
		{SrcTemplate: "core/builds/Dockerfile.tmpl", DstPath: "[[.ServiceName]]/builds/Dockerfile"},
		{SrcTemplate: "core/config/config.local.example.yaml.tmpl", DstPath: "[[.ServiceName]]/config/config.local.example.yaml"},
		{SrcTemplate: "core/env.example.tmpl", DstPath: "[[.ServiceName]]/.env.example"},
		{SrcTemplate: "core/docker-compose.yml.tmpl", DstPath: "[[.ServiceName]]/docker-compose.yml"},

		// ─── Framework: Echo ────────────────────────────────────────────────
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/router.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/router.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/api/v1/v1_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/api/v1/v1_handler.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/api/v1/user_v1_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/api/v1/user_v1_handler.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/health_check/health_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/health_check/health_handler.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/middleware/middleware.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/middleware.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/middleware/cors.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/cors.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/middleware/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/logger.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/middleware/recover.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/recover.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/middleware/api_key.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/api_key.go",
			Condition:   isEcho,
		},
		{
			SrcTemplate: "frameworks/echo/internal/deliveries/http/middleware/jwt.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/jwt.go",
			Condition:   func(c *Config) bool { return isEcho(c) && c.UseJWT },
		},

		// ─── Framework: Gin ─────────────────────────────────────────────────
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/router.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/router.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/api/v1/v1_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/api/v1/v1_handler.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/api/v1/user_v1_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/api/v1/user_v1_handler.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/health_check/health_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/health_check/health_handler.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/middleware/middleware.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/middleware.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/middleware/cors.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/cors.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/middleware/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/logger.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/middleware/recover.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/recover.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/middleware/api_key.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/api_key.go",
			Condition:   isGin,
		},
		{
			SrcTemplate: "frameworks/gin/internal/deliveries/http/middleware/jwt.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/jwt.go",
			Condition:   func(c *Config) bool { return isGin(c) && c.UseJWT },
		},

		// ─── Framework: Fiber ───────────────────────────────────────────────
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/router.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/router.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/api/v1/v1_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/api/v1/v1_handler.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/api/v1/user_v1_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/api/v1/user_v1_handler.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/health_check/health_handler.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/health_check/health_handler.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/middleware/middleware.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/middleware.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/middleware/cors.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/cors.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/middleware/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/logger.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/middleware/recover.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/recover.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/middleware/api_key.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/api_key.go",
			Condition:   isFiber,
		},
		{
			SrcTemplate: "frameworks/fiber/internal/deliveries/http/middleware/jwt.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/deliveries/http/middleware/jwt.go",
			Condition:   func(c *Config) bool { return isFiber(c) && c.UseJWT },
		},

		// ─── Database: PostgreSQL + GORM ────────────────────────────────────
		{
			SrcTemplate: "databases/postgres_gorm/internal/pkg/database/postgres.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/postgres.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_gorm/internal/repository/pgsql/main_pgsql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/main_pgsql_repository.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_gorm/internal/repository/pgsql/health_pgsql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/health_pgsql_repository.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_gorm/internal/repository/pgsql/user_pgsql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/user_pgsql_repository.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_gorm/internal/repository/pgsql/transaction_pgsql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/transaction_pgsql_repository.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGORM },
		},

		// ─── Database: PostgreSQL + Raw (pgx) ───────────────────────────────
		{
			SrcTemplate: "databases/postgres_raw/internal/pkg/database/postgres.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/postgres.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_raw/internal/repository/pgsql/main_pgsql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/main_pgsql_repository.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_raw/internal/repository/pgsql/health_pgsql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/health_pgsql_repository.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_raw/internal/repository/pgsql/user_pgsql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/user_pgsql_repository.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && !c.UseGORM },
		},

		// ─── Database: MySQL + GORM ─────────────────────────────────────────
		{
			SrcTemplate: "databases/mysql_gorm/internal/pkg/database/mysql.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/mysql.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_gorm/internal/repository/mysql/main_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/main_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_gorm/internal/repository/mysql/health_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/health_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_gorm/internal/repository/mysql/user_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/user_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGORM },
		},

		// ─── Database: MySQL + Raw (sqlx) ───────────────────────────────────
		{
			SrcTemplate: "databases/mysql_raw/internal/pkg/database/mysql.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/mysql.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_raw/internal/repository/mysql/main_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/main_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_raw/internal/repository/mysql/health_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/health_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_raw/internal/repository/mysql/user_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/user_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && !c.UseGORM },
		},

		// ─── Database: MongoDB ──────────────────────────────────────────────
		{
			SrcTemplate: "databases/mongodb/internal/pkg/database/mongo.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/mongo.go",
			Condition:   func(c *Config) bool { return c.HasMongoDB },
		},
		{
			SrcTemplate: "databases/mongodb/internal/repository/mongo/main_mongo_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mongo/main_mongo_repository.go",
			Condition:   func(c *Config) bool { return c.HasMongoDB },
		},
		{
			SrcTemplate: "databases/mongodb/internal/repository/mongo/health_mongo_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mongo/health_mongo_repository.go",
			Condition:   func(c *Config) bool { return c.HasMongoDB },
		},
		{
			SrcTemplate: "databases/mongodb/internal/repository/mongo/user_mongo_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mongo/user_mongo_repository.go",
			Condition:   func(c *Config) bool { return c.HasMongoDB },
		},

		// ─── Database: Redis ────────────────────────────────────────────────
		{
			SrcTemplate: "databases/redis/internal/pkg/database/redis.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/redis.go",
			Condition:   func(c *Config) bool { return c.HasRedis },
		},
		{
			SrcTemplate: "databases/redis/internal/repository/cache/main_cache_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/cache/main_cache_repository.go",
			Condition:   func(c *Config) bool { return c.HasRedis },
		},

		// ─── database.go (aggregator) — always if any DB ────────────────────
		{
			SrcTemplate: "core/internal/pkg/database/database.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/database.go",
			Condition:   func(c *Config) bool { return c.HasAnyDB() },
		},

		// ─── Logger ─────────────────────────────────────────────────────────
		{
			SrcTemplate: "loggers/zap/internal/pkg/logger/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/logger.go",
			Condition:   func(c *Config) bool { return c.Logger == "zap" },
		},
		{
			SrcTemplate: "loggers/zerolog/internal/pkg/logger/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/logger.go",
			Condition:   func(c *Config) bool { return c.Logger == "zerolog" },
		},
		{
			SrcTemplate: "loggers/logrus/internal/pkg/logger/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/logger.go",
			Condition:   func(c *Config) bool { return c.Logger == "logrus" },
		},
		{
			SrcTemplate: "loggers/slog/internal/pkg/logger/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/logger.go",
			Condition:   func(c *Config) bool { return c.Logger == "slog" },
		},

		// ─── Tracing: OpenTelemetry ─────────────────────────────────────────
		{
			SrcTemplate: "tracing/otel/internal/pkg/tracer/tracer.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/tracer/tracer.go",
			Condition:   func(c *Config) bool { return c.UseOTel },
		},

		// ─── Auth: JWT ──────────────────────────────────────────────────────
		{
			SrcTemplate: "auth/jwt/internal/pkg/jwtc/jwtc.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/jwtc/jwtc.go",
			Condition:   func(c *Config) bool { return c.UseJWT },
		},

		// ─── Swagger ────────────────────────────────────────────────────────
		{
			SrcTemplate: "swagger/docs/docs.go.tmpl",
			DstPath:     "[[.ServiceName]]/docs/docs.go",
			Condition:   func(c *Config) bool { return c.UseSwagger },
		},

		// ─── Migration: Goose ───────────────────────────────────────────────
		{
			SrcTemplate: "migrations/goose/postgre/001_init.sql.tmpl",
			DstPath:     "[[.ServiceName]]/migration/db/postgre/001_init.sql",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGoose() },
		},
		{
			SrcTemplate: "migrations/goose/mysql/001_init.sql.tmpl",
			DstPath:     "[[.ServiceName]]/migration/db/mysql/001_init.sql",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGoose() },
		},
		{
			SrcTemplate: "migrations/golang_migrate/postgre/000001_init.up.sql.tmpl",
			DstPath:     "[[.ServiceName]]/migration/db/postgre/000001_init.up.sql",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGolangMigrate() },
		},
		{
			SrcTemplate: "migrations/golang_migrate/postgre/000001_init.down.sql.tmpl",
			DstPath:     "[[.ServiceName]]/migration/db/postgre/000001_init.down.sql",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGolangMigrate() },
		},
		{
			SrcTemplate: "migrations/golang_migrate/mysql/000001_init.up.sql.tmpl",
			DstPath:     "[[.ServiceName]]/migration/db/mysql/000001_init.up.sql",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGolangMigrate() },
		},
		{
			SrcTemplate: "migrations/golang_migrate/mysql/000001_init.down.sql.tmpl",
			DstPath:     "[[.ServiceName]]/migration/db/mysql/000001_init.down.sql",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGolangMigrate() },
		},
	}
}

// ─── Condition helpers ───────────────────────────────────────────────────────

func isEcho(c *Config) bool  { return c.Framework == "echo" }
func isGin(c *Config) bool   { return c.Framework == "gin" }
func isFiber(c *Config) bool { return c.Framework == "fiber" }
