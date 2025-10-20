package util

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "util",
	}

	cmd.AddCommand(
		dateCommand(),
		randomCommand(),
	)

	return cmd
}
