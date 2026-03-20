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
		{SrcTemplate: "core/internal/models/token_model.go.tmpl", DstPath: "[[.ServiceName]]/internal/models/token_model.go"},
		{SrcTemplate: "core/internal/pkg/boolc/boolc.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/boolc/boolc.go"},
		{SrcTemplate: "core/internal/pkg/errorc/model.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/errorc/model.go"},
		{SrcTemplate: "core/internal/pkg/formatter/phone_number.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/formatter/phone_number.go"},
		{SrcTemplate: "core/internal/pkg/generator/account_number.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/account_number.go"},
		{SrcTemplate: "core/internal/pkg/generator/account_number_test.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/account_number_test.go"},
		{SrcTemplate: "core/internal/pkg/generator/hash.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/hash.go"},
		{SrcTemplate: "core/internal/pkg/generator/hash_test.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/hash_test.go"},
		{SrcTemplate: "core/internal/pkg/generator/jwt.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/jwt.go"},
		{SrcTemplate: "core/internal/pkg/generator/jwt_test.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/jwt_test.go"},
		{SrcTemplate: "core/internal/pkg/generator/token.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/token.go"},
		{SrcTemplate: "core/internal/pkg/generator/token_test.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/generator/token_test.go"},
		{SrcTemplate: "core/internal/pkg/numberc/numberc.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/numberc/numberc.go"},
		{SrcTemplate: "core/internal/pkg/response/error_response.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/response/error_response.go"},
		{SrcTemplate: "core/internal/pkg/response/success_response.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/response/success_response.go"},
		{SrcTemplate: "core/internal/pkg/stringc/stringc.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/stringc/stringc.go"},
		{SrcTemplate: "core/internal/pkg/validator/account_number.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/account_number.go"},
		{SrcTemplate: "core/internal/pkg/validator/account_number_test.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/account_number_test.go"},
		{SrcTemplate: "core/internal/pkg/validator/hash.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/hash.go"},
		{SrcTemplate: "core/internal/pkg/validator/jwt.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/jwt.go"},
		{SrcTemplate: "core/internal/pkg/validator/jwt_test.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/jwt_test.go"},
		{SrcTemplate: "core/internal/pkg/validator/password.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/password.go"},
		{SrcTemplate: "core/internal/pkg/validator/phone_number.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/phone_number.go"},
		{SrcTemplate: "core/internal/pkg/validator/rest.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/rest.go"},
		{SrcTemplate: "core/internal/pkg/validator/validation_mapper.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/validator/validation_mapper.go"},
		{SrcTemplate: "core/internal/pkg/graceful/graceful.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/graceful.go"},
		{SrcTemplate: "core/internal/pkg/graceful/process.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/process.go"},
		{SrcTemplate: "core/internal/pkg/graceful/options.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/options.go"},
		{SrcTemplate: "core/internal/pkg/graceful/logger.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/logger.go"},
		{SrcTemplate: "core/internal/pkg/graceful/logger_adapter.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/logger_adapter.go"},
		{SrcTemplate: "core/internal/pkg/graceful/func_process.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/func_process.go"},
		{SrcTemplate: "core/internal/pkg/graceful/http_process.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/graceful/[[.Framework]]_process.go"},
		{SrcTemplate: "core/internal/pkg/errorc/errorc.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/errorc/errorc.go"},
		{SrcTemplate: "core/internal/pkg/formatter/formatter.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/formatter/formatter.go"},
		{SrcTemplate: "core/internal/pkg/logger/logger.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/logger/logger.go"},
		{SrcTemplate: "core/internal/pkg/logger/context.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/logger/context.go"},
		{SrcTemplate: "core/internal/pkg/logger/error.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/logger/error.go"},
		{SrcTemplate: "core/internal/pkg/logger/masking.go.tmpl", DstPath: "[[.ServiceName]]/internal/pkg/logger/masking.go"},
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

		// ─── Database: PostgreSQL Connections ───────────────────────────────
		{
			SrcTemplate: "core/internal/pkg/database/postgre_database.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/postgre_database.go",
			Condition:   func(c *Config) bool { return c.HasPostgres },
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

		// ─── Database: PostgreSQL + Raw (pgx) Repositories ──────────────────
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
		{
			SrcTemplate: "databases/postgres_raw/internal/repository/pgsql/user_pg_query.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/user_pg_query.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/postgres_gorm/internal/repository/pgsql/user_pg_query.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/pgsql/user_pg_query.go",
			Condition:   func(c *Config) bool { return c.HasPostgres && c.UseGORM },
		},

		// ─── Database: MySQL Connections ────────────────────────────────────
		{
			SrcTemplate: "core/internal/pkg/database/mysql_database.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/mysql_database.go",
			Condition:   func(c *Config) bool { return c.HasMySQL },
		},
		{
			SrcTemplate: "databases/mysql_gorm/internal/repository/mysql/main_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/main_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_gorm/internal/repository/mysql/user_msql_query.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/user_msql_query.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && c.UseGORM },
		},

		// ─── Database: MySQL + Raw (sqlx) Repositories ──────────────────────
		{
			SrcTemplate: "databases/mysql_raw/internal/repository/mysql/main_mysql_repository.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/main_mysql_repository.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && !c.UseGORM },
		},
		{
			SrcTemplate: "databases/mysql_raw/internal/repository/mysql/user_msql_query.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mysql/user_msql_query.go",
			Condition:   func(c *Config) bool { return c.HasMySQL && !c.UseGORM },
		},

		// ─── Database: MongoDB ──────────────────────────────────────────────
		{
			SrcTemplate: "core/internal/pkg/database/mongo_database.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/mongo_database.go",
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
		{
			SrcTemplate: "databases/mongodb/internal/repository/mongo/user_mongo_query.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/repository/mongo/user_mongo_query.go",
			Condition:   func(c *Config) bool { return c.HasMongoDB },
		},

		// ─── Database: Redis ────────────────────────────────────────────────
		{
			SrcTemplate: "core/internal/pkg/database/redis_database.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/database/redis_database.go",
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
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/zap.go",
			Condition:   func(c *Config) bool { return c.Logger == "zap" },
		},
		{
			SrcTemplate: "loggers/zap/internal/pkg/logger/logger_bench_test.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/zap_bench_test.go",
			Condition:   func(c *Config) bool { return c.Logger == "zap" },
		},
		{
			SrcTemplate: "loggers/zerolog/internal/pkg/logger/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/zerolog.go",
			Condition:   func(c *Config) bool { return c.Logger == "zerolog" },
		},
		{
			SrcTemplate: "loggers/logrus/internal/pkg/logger/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/logrus.go",
			Condition:   func(c *Config) bool { return c.Logger == "logrus" },
		},
		{
			SrcTemplate: "loggers/slog/internal/pkg/logger/logger.go.tmpl",
			DstPath:     "[[.ServiceName]]/internal/pkg/logger/slog.go",
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
		{
			SrcTemplate: "swagger/docs/api-docs.html.tmpl",
			DstPath:     "[[.ServiceName]]/docs/api-docs.html",
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
