package cli

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "mcphome",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start MCP server",
		RunE:  handleStart,
	}

	startCmd.Flags().IP("host", net.ParseIP("127.0.0.1"), "Host to run the MCP server on")
	startCmd.Flags().Uint16("port", 8080, "Port to run the MCP server on")

	var pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Install MCP plugins",
		RunE:  handlePull,
	}

	rootCmd.AddCommand(startCmd, pullCmd)

	return rootCmd
}

func handleStart(cmd *cobra.Command, args []string) error {
	host, errHost := cmd.Flags().GetIP("host")
	port, errPort := cmd.Flags().GetUint16("port")

	if errHost != nil {
		return fmt.Errorf("failed to get host: %w", errHost)
	}

	if errPort != nil {
		return fmt.Errorf("failed to get port: %w", errPort)
	}

	fmt.Printf("Starting MCP server on %s:%d...\n", host.String(), port)
	// Add logic to start the server here

	return nil
}

func handlePull(cmd *cobra.Command, args []string) error {
	fmt.Println("Pulling MCP plugins...")
	// Add logic to pull plugins here

	return nil
}
