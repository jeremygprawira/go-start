package generator

import (
	"fmt"
	"strings"
)

// buildGoMod dynamically constructs a go.mod file based on selections.
func buildGoMod(cfg *Config) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("module %s\n\ngo 1.23.4\n\nrequire (\n", cfg.ModuleName))

	// --- Framework ---
	switch cfg.Framework {
	case "echo":
		sb.WriteString("\tgithub.com/labstack/echo/v4 v4.13.4\n")
	case "gin":
		sb.WriteString("\tgithub.com/gin-gonic/gin v1.10.0\n")
	case "fiber":
		sb.WriteString("\tgithub.com/gofiber/fiber/v2 v2.52.5\n")
	}

	// --- Database ---
	if cfg.HasPostgres {
		if cfg.UseGORM {
			sb.WriteString("\tgorm.io/gorm v1.30.0\n")
			sb.WriteString("\tgorm.io/driver/postgres v1.6.0\n")
		} else {
			sb.WriteString("\tgithub.com/jackc/pgx/v5 v5.6.0\n")
		}
	}
	if cfg.HasMySQL {
		if cfg.UseGORM {
			sb.WriteString("\tgorm.io/gorm v1.30.0\n")
			sb.WriteString("\tgorm.io/driver/mysql v1.5.7\n")
		} else {
			sb.WriteString("\tgithub.com/jmoiron/sqlx v1.4.0\n")
			sb.WriteString("\tgithub.com/go-sql-driver/mysql v1.8.1\n")
		}
	}
	if cfg.HasMongoDB {
		sb.WriteString("\tgo.mongodb.org/mongo-driver/v2 v2.1.0\n")
	}
	if cfg.HasRedis {
		sb.WriteString("\tgithub.com/redis/go-redis/v9 v9.7.3\n")
	}

	// --- Migration ---
	if cfg.UseGoose() {
		sb.WriteString("\tgithub.com/pressly/goose/v3 v3.24.2\n")
	}
	if cfg.UseGolangMigrate() {
		sb.WriteString("\tgithub.com/golang-migrate/migrate/v4 v4.18.2\n")
	}

	// --- Logger ---
	switch cfg.Logger {
	case "zap":
		sb.WriteString("\tgo.uber.org/zap v1.27.0\n")
	case "zerolog":
		sb.WriteString("\tgithub.com/rs/zerolog v1.34.0\n")
	case "logrus":
		sb.WriteString("\tgithub.com/sirupsen/logrus v1.9.3\n")
	// slog is stdlib, no import needed
	}

	// --- Tracing ---
	if cfg.UseOTel {
		sb.WriteString("\tgo.opentelemetry.io/otel v1.35.0\n")
		sb.WriteString("\tgo.opentelemetry.io/otel/trace v1.35.0\n")
		sb.WriteString("\tgo.opentelemetry.io/otel/sdk v1.35.0\n")
		sb.WriteString("\tgo.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.35.0\n")
	}

	// --- Auth ---
	if cfg.UseJWT {
		sb.WriteString("\tgithub.com/golang-jwt/jwt/v5 v5.2.2\n")
	}

	// --- Swagger ---
	if cfg.UseSwagger {
		switch cfg.Framework {
		case "echo":
			sb.WriteString("\tgithub.com/swaggo/echo-swagger v1.4.1\n")
		case "gin":
			sb.WriteString("\tgithub.com/swaggo/gin-swagger v1.6.0\n")
		case "fiber":
			sb.WriteString("\tgithub.com/gofiber/swagger v1.1.1\n")
		}
		sb.WriteString("\tgithub.com/swaggo/swag v1.8.12\n")
	}

	// --- Always included ---
	sb.WriteString("\tgithub.com/spf13/viper v1.20.1\n")
	sb.WriteString("\tgolang.org/x/crypto v0.38.0\n")
	sb.WriteString("\tgithub.com/go-playground/validator/v10 v10.27.0\n")
	sb.WriteString("\tgithub.com/google/uuid v1.6.0\n")
	sb.WriteString(")\n")

	return sb.String()
}
