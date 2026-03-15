package cmd

import (
	"fmt"
	"go-start/internal/generator"
	"go-start/internal/ui"
	"os"
	"path/filepath"
	"runtime"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Scaffold a new Go project interactively",
	Long:  `Interactively scaffold a production-ready Go project with clean architecture.`,
	Run:   runNew,
}

func runNew(cmd *cobra.Command, args []string) {
	ui.PrintBanner()

	cfg := &generator.Config{}

	// ─── 1. Service Name ────────────────────────────────────────────────────
	name, err := promptText("Service Name", "my-service", func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("service name cannot be empty")
		}
		if _, err := os.Stat(input); !os.IsNotExist(err) {
			return fmt.Errorf("directory %q already exists", input)
		}
		return nil
	})
	exitOnErr(err)
	cfg.ServiceName = name

	// ─── 2. Go Module Name ──────────────────────────────────────────────────
	module, err := promptText("Go Module Name", fmt.Sprintf("github.com/yourname/%s", name), func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("module name cannot be empty")
		}
		return nil
	})
	exitOnErr(err)
	cfg.ModuleName = module

	// ─── 3. Framework ───────────────────────────────────────────────────────
	_, framework, err := promptSelect("HTTP Framework", []string{"echo", "gin", "fiber"})
	exitOnErr(err)
	cfg.Framework = framework

	// ─── 4. Databases (multi-select) ────────────────────────────────────────
	dbChoices := []string{"PostgreSQL", "MySQL", "MongoDB", "Redis"}
	selectedDBs, err := ui.PromptMultiSelect("Select Databases", dbChoices)
	exitOnErr(err)

	for _, db := range selectedDBs {
		switch db {
		case "PostgreSQL":
			cfg.HasPostgres = true
		case "MySQL":
			cfg.HasMySQL = true
		case "MongoDB":
			cfg.HasMongoDB = true
		case "Redis":
			cfg.HasRedis = true
		}
	}

	// ─── 5. ORM (only for SQL DBs) ──────────────────────────────────────────
	if cfg.HasPostgres || cfg.HasMySQL {
		_, orm, err := promptSelect("Database Access Layer", []string{"GORM (ORM)", "Raw SQL (pgx / sqlx)"})
		exitOnErr(err)
		cfg.UseGORM = orm == "GORM (ORM)"
	}

	// ─── 6. Migration Tool (only for SQL DBs) ───────────────────────────────
	if cfg.HasPostgres || cfg.HasMySQL {
		_, migration, err := promptSelect("Migration Tool", []string{"goose", "golang-migrate", "none"})
		exitOnErr(err)
		cfg.MigrationTool = migration
	} else {
		cfg.MigrationTool = "none"
	}

	// ─── 7. Logger ──────────────────────────────────────────────────────────
	_, logger, err := promptSelect("Logger", []string{"zap", "zerolog", "logrus", "slog (stdlib)"})
	exitOnErr(err)
	if logger == "slog (stdlib)" {
		logger = "slog"
	}
	cfg.Logger = logger

	// ─── 8. Tracing ─────────────────────────────────────────────────────────
	_, tracing, err := promptSelect("Distributed Tracing", []string{"OpenTelemetry (OTLP/gRPC)", "none"})
	exitOnErr(err)
	cfg.UseOTel = tracing == "OpenTelemetry (OTLP/gRPC)"

	// ─── 9. Auth ────────────────────────────────────────────────────────────
	_, auth, err := promptSelect("Authentication", []string{"JWT", "none"})
	exitOnErr(err)
	cfg.UseJWT = auth == "JWT"

	// ─── 10. Swagger ────────────────────────────────────────────────────────
	_, swagger, err := promptSelect("API Documentation", []string{"Swagger (swaggo)", "none"})
	exitOnErr(err)
	cfg.UseSwagger = swagger == "Swagger (swaggo)"

	// ─── Summary ────────────────────────────────────────────────────────────
	fmt.Println()
	ui.PrintSummary(cfg)
	fmt.Println()

	confirmed, err := promptConfirm("Generate project?")
	exitOnErr(err)
	if !confirmed {
		ui.PrintInfo("Aborted.")
		os.Exit(0)
	}

	// ─── Generate ───────────────────────────────────────────────────────────
	templatesDir := findTemplatesDir()

	s := ui.NewSpinner("Generating project...")
	s.Start()
	err = generator.Generate(cfg, templatesDir)
	s.Stop()

	if err != nil {
		ui.PrintError("Failed to generate project: " + err.Error())
		os.Exit(1)
	}

	ui.PrintSuccess("Project generated successfully!")
	fmt.Println()
	fmt.Printf("  📁 %s/\n", cfg.ServiceName)
	ui.PrintTree(cfg.ServiceName, "  ")
	fmt.Println()
	fmt.Println("  Next steps:")
	fmt.Printf("  %s %s\n", ui.Cyan("cd"), cfg.ServiceName)
	fmt.Printf("  %s\n", ui.Cyan("docker compose up -d"))
	fmt.Printf("  %s\n", ui.Cyan("make run"))
	fmt.Println()
	fmt.Println("  Happy coding! 🚀")
}

// ─── Prompt Helpers ─────────────────────────────────────────────────────────

func promptText(label, defaultVal string, validate func(string) error) (string, error) {
	p := promptui.Prompt{
		Label:    label,
		Default:  defaultVal,
		Validate: validate,
	}
	return p.Run()
}

func promptSelect(label string, items []string) (int, string, error) {
	p := promptui.Select{
		Label: label,
		Items: items,
		Size:  8,
	}
	return p.Run()
}

func promptConfirm(label string) (bool, error) {
	p := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	result, err := p.Run()
	if err != nil {
		// promptui returns err on 'n' answer
		if err == promptui.ErrAbort {
			return false, nil
		}
		return false, err
	}
	return result == "y" || result == "Y", nil
}

func exitOnErr(err error) {
	if err != nil {
		if err == promptui.ErrInterrupt {
			ui.PrintError("Interrupted.")
		} else {
			ui.PrintError(err.Error())
		}
		os.Exit(1)
	}
}

// findTemplatesDir locates the templates/ directory relative to the binary.
func findTemplatesDir() string {
	// 1. Next to the binary
	exe, err := os.Executable()
	if err == nil {
		candidate := filepath.Join(filepath.Dir(exe), "templates")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	// 2. Relative to source (for `go run`)
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		// filename is cmd/new.go, templates is at root
		candidate := filepath.Join(filepath.Dir(filepath.Dir(filename)), "templates")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	// 3. Current dir fallback
	return "templates"
}
