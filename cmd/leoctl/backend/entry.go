package backend

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "backend",
	}

	cmd.AddCommand(
		snippetCommand(),
	)

	return cmd
}
