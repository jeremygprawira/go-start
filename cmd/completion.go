// cmd/completion.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
Explanation:
Generate shell completion scripts
*/

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion scripts",
	Long: `To load completions:

Bash:
	source <(go-start completion bash)

Zsh:
	echo "autoload -U compinit; compinit" >> ~/.zshrc
	go-start completion zsh > "${fpath[1]}/_go-app-boilerplate"

Fish:
	go-start completion fish | source

PowerShell:
	go-start completion powershell | Out-String | Invoke-Expression
`,
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("requires exactly 1 argument: bash, zsh, fish, or powershell")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported shell: %s\n", args[0])
		}
	},
}
