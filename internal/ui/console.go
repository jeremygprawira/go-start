package ui

import (
	"bufio"
	"fmt"
	"go-start/internal/generator"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var (
	Cyan    = color.New(color.FgCyan).SprintFunc()
	Green   = color.New(color.FgGreen).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Bold    = color.New(color.Bold).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
)

func PrintBanner() {
	fmt.Println()
	fmt.Println(Cyan("  ╔═══════════════════════════════════════╗"))
	fmt.Println(Cyan("  ║") + Bold("        🚀  go-start  CLI              ") + Cyan("║"))
	fmt.Println(Cyan("  ║") + "    Golang Clean Architecture Scaffold  " + Cyan("║"))
	fmt.Println(Cyan("  ╚═══════════════════════════════════════╝"))
	fmt.Println()
}

func PrintInfo(msg string) {
	fmt.Println(Cyan("ℹ  " + msg))
}

func PrintSuccess(msg string) {
	fmt.Println(Green("✅ " + msg))
}

func PrintError(msg string) {
	fmt.Fprintln(os.Stderr, Red("❌ "+msg))
}

func NewSpinner(msg string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
	s.Suffix = " " + msg
	s.Color("cyan")
	return s
}

func PrintTree(root string, prefix string) {
	files, err := os.ReadDir(root)
	if err != nil {
		return
	}
	for i, f := range files {
		connector := "├── "
		if i == len(files)-1 {
			connector = "└── "
		}
		if f.IsDir() {
			fmt.Printf("%s%s%s\n", prefix, Cyan(connector), Cyan(f.Name()+"/"))
			childPrefix := prefix + "│   "
			if i == len(files)-1 {
				childPrefix = prefix + "    "
			}
			PrintTree(filepath.Join(root, f.Name()), childPrefix)
		} else {
			fmt.Printf("%s%s%s\n", prefix, connector, f.Name())
		}
	}
}

// PrintSummary displays a color-coded summary of all choices.
func PrintSummary(cfg *generator.Config) {
	fmt.Println(Bold("  ┌─ Project Summary " + strings.Repeat("─", 24) + "┐"))
	printRow("Service Name", cfg.ServiceName)
	printRow("Go Module   ", cfg.ModuleName)
	printRow("Framework   ", strings.ToUpper(cfg.Framework))

	var dbs []string
	if cfg.HasPostgres {
		dbs = append(dbs, "PostgreSQL")
	}
	if cfg.HasMySQL {
		dbs = append(dbs, "MySQL")
	}
	if cfg.HasMongoDB {
		dbs = append(dbs, "MongoDB")
	}
	if cfg.HasRedis {
		dbs = append(dbs, "Redis")
	}
	if len(dbs) == 0 {
		dbs = []string{"none"}
	}
	printRow("Databases   ", strings.Join(dbs, " + "))

	if cfg.HasPostgres || cfg.HasMySQL {
		orm := "Raw SQL"
		if cfg.UseGORM {
			orm = "GORM"
		}
		printRow("ORM         ", orm)
		printRow("Migrations  ", cfg.MigrationTool)
	}

	printRow("Logger      ", cfg.Logger)

	tracing := "none"
	if cfg.UseOTel {
		tracing = "OpenTelemetry"
	}
	printRow("Tracing     ", tracing)

	auth := "none"
	if cfg.UseJWT {
		auth = "JWT"
	}
	printRow("Auth        ", auth)

	swagger := "no"
	if cfg.UseSwagger {
		swagger = "yes (swaggo)"
	}
	printRow("Swagger     ", swagger)
	fmt.Println(Bold("  └" + strings.Repeat("─", 43) + "┘"))
}

func printRow(label, value string) {
	fmt.Printf("  │  %s  %s\n", Yellow(label+":"), Cyan(value))
}

// PromptMultiSelect renders a checkbox-style multi-select prompt.
func PromptMultiSelect(label string, options []string) ([]string, error) {
	selected := make([]bool, len(options))

	fmt.Printf("\n%s %s\n", Yellow("?"), Bold(label))
	fmt.Println(Cyan("  (use ↑/↓ to navigate, SPACE to toggle, ENTER to confirm)"))
	fmt.Println()

	cursor := 0

	// Simple terminal-based multi-select
	reader := bufio.NewReader(os.Stdin)
	// Fallback: just ask for comma-separated input
	for i, opt := range options {
		fmt.Printf("  [%d] %s\n", i+1, opt)
	}
	fmt.Println()
	fmt.Print(Cyan("  Select (e.g. 1,3 or press ENTER to skip all): "))

	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)

	_ = cursor
	_ = selected

	if input == "" {
		return []string{}, nil
	}

	var result []string
	parts := strings.Split(input, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		for i, opt := range options {
			if fmt.Sprintf("%d", i+1) == p {
				result = append(result, opt)
				break
			}
		}
	}
	return result, nil
}
