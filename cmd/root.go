package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-start",
	Short: "A CLI tool to scaffold production-ready Go projects",
	Long: `go-start scaffolds Golang projects with clean architecture,
giving you full flexibility over your framework, database,
logger, tracing, auth, and more.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
}
