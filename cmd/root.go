package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-start",
	Short: "A CLI tool to start your Go Programming Language or Golang projects",
	Long:  `Start your Go projects using Echo, Gin, or Fiber with clean architecture, GORM/raw SQL, JWT, and logger presets.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scaffoldCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(initCmd)
}
