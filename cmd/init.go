package cmd

import (
	"fmt"
	"go-start/internal/starter"
	"go-start/internal/ui"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

/*
Trigger the interactive process to start a new Go project.
*/

var initCmd = &cobra.Command{
	Use:   "now",
	Short: "Create a new Go project interactively",
	Run: func(cmd *cobra.Command, args []string) {
		// Project Name
		prompt := promptui.Prompt{
			Label: "Project Name",
			Validate: func(input string) error {
				if len(input) == 0 {
					return fmt.Errorf("name cannot be empty")
				}
				return nil
			},
		}
		name, err := prompt.Run()
		if err != nil {
			ui.PrintError("Aborted.")
			os.Exit(1)
		}

		// Framework
		frameworkPrompt := promptui.Select{
			Label: "Select Framework",
			Items: []string{"echo", "gin - coming soon!", "fiber - coming soon!"},
		}
		_, framework, err := frameworkPrompt.Run()
		if err != nil {
			ui.PrintError("Aborted.")
			os.Exit(1)
		}

		// Validate options based on template readiness
		if framework == "gin - coming soon!" || framework == "fiber - coming soon!" {
			ui.PrintError("Selected framework is not yet supported. Please choose 'echo'.")
			os.Exit(1)
		}

		// Database
		dbPrompt := promptui.Select{
			Label: "Select Database Engine",
			Items: []string{"gorm", "raw"},
		}
		_, db, err := dbPrompt.Run()
		if err != nil {
			ui.PrintError("Aborted.")
			os.Exit(1)
		}

		// Logger
		logPrompt := promptui.Select{
			Label: "Select Logger",
			Items: []string{"zap", "zerolog", "logrus"},
		}
		_, logType, err := logPrompt.Run()
		if err != nil {
			ui.PrintError("Aborted.")
			os.Exit(1)
		}

		// Trigger UI Implementation of Spinner
		s := ui.NewSpinner("Generating project...")
		s.Start()
		err = starter.Generate(name, framework, db, logType)
		s.Stop()

		if err != nil {
			ui.PrintError("Failed to create project: " + err.Error())
			os.Exit(1)
		}

		// Information
		ui.PrintSuccess("Project generated successfully!")
		fmt.Println(ui.Cyan("\n📁 " + name + "/"))
		ui.PrintTree(name, "")
		fmt.Println("\nNext steps:")
		fmt.Println(ui.Cyan("  cd"), name)
		fmt.Println(ui.Cyan("  docker-compose up --build"))
		fmt.Println("\nHappy coding! 🚀")
	},
}
