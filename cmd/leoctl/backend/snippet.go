package backend

import (
	"leo/internal/biz/snippet"

	"github.com/spf13/cobra"
)

func snippetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "snippet",
	}

	cmd.AddCommand(
		&cobra.Command{
			Use: "add",
			Run: func(_ *cobra.Command, args []string) {
				if len(args) > 1 {
					snippet.Add(args[0], args[1])
				}
			},
		},
		&cobra.Command{
			Use: "rm",
			Run: func(_ *cobra.Command, args []string) {
				if len(args) > 1 {
					snippet.Rm(args[0], args[1])
				}
			},
		},
	)

	return cmd
}
