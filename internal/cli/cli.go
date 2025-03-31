package cli

import (
	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "mcphome",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		PersistentPreRunE: HandlePreRun,
	}

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start MCP server",
		RunE:  HandleStart,
	}

	var pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Install MCP plugins",
		RunE:  HandlePull,
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List installed MCP plugins",
		RunE:  HandleList,
	}

	rootCmd.AddCommand(startCmd, pullCmd, listCmd)

	return rootCmd
}
