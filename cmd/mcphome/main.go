package main

import (
	"context"

	"github.com/mcphome/mcphome/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cli.NewCLI().ExecuteContext(context.Background()))
}
