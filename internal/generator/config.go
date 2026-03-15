package generator

// Config holds all user choices for the project to be generated.
type Config struct {
	// Basic Info
	ServiceName string // directory name + docker container name
	ModuleName  string // Go module path, e.g. github.com/yourname/myservice

	// Framework
	Framework string // "echo" | "gin" | "fiber"

	// Databases (multi-select)
	HasPostgres bool
	HasMySQL    bool
	HasMongoDB  bool
	HasRedis    bool

	// ORM (only relevant if HasPostgres || HasMySQL)
	UseGORM   bool // true = GORM, false = raw (pgx/sqlx)

	// Migration tool (only relevant if SQL DB)
	MigrationTool string // "goose" | "golang-migrate" | "none"

	// Logger
	Logger string // "zap" | "zerolog" | "logrus" | "slog"

	// Tracing
	UseOTel bool // OpenTelemetry

	// Auth
	UseJWT bool

	// Swagger
	UseSwagger bool
}

// HasSQLDB returns true if any SQL database is selected.
func (c *Config) HasSQLDB() bool {
	return c.HasPostgres || c.HasMySQL
}

// HasAnyDB returns true if any database is selected.
func (c *Config) HasAnyDB() bool {
	return c.HasPostgres || c.HasMySQL || c.HasMongoDB || c.HasRedis
}

// NeedsMigration returns true if migration files should be generated.
func (c *Config) NeedsMigration() bool {
	return c.HasSQLDB() && c.MigrationTool != "none"
}

// UseGoose returns true if goose is the migration tool.
func (c *Config) UseGoose() bool {
	return c.MigrationTool == "goose"
}

// UseGolangMigrate returns true if golang-migrate is the migration tool.
func (c *Config) UseGolangMigrate() bool {
	return c.MigrationTool == "golang-migrate"
}
