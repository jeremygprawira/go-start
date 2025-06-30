package cmd

import (
	"fmt"
	"os"

	"go-start/internal/starter"

	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new project with clean architecture and selected framework",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		framework, _ := cmd.Flags().GetString("framework")
		db, _ := cmd.Flags().GetString("db")
		logType, _ := cmd.Flags().GetString("log")

		if name == "" {
			fmt.Println("❌ Project name is required (--name)")
			os.Exit(1)
		}

		err := starter.Generate(name, framework, db, logType)
		if err != nil {
			fmt.Println("❌ Failed:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Project generated successfully")
	},
}

func init() {
	scaffoldCmd.Flags().StringP("name", "n", "", "Project name (required)")
	scaffoldCmd.Flags().StringP("framework", "f", "echo", "Framework to use: echo, gin, fiber")
	scaffoldCmd.Flags().StringP("db", "d", "gorm", "Database engine: gorm or raw")
	scaffoldCmd.Flags().StringP("log", "l", "zap", "Logger to use: zap, zerolog, or logrus")
}
