package main

import (
	"leo/cmd/leoctl/backend"
	"leo/cmd/leoctl/ide"
	"leo/cmd/leoctl/snippet"
	"leo/cmd/leoctl/util"
	"leo/internal/pb"

	aw "github.com/deanishe/awgo"
	"github.com/spf13/cobra"
)

var (
	_rootCmd = &cobra.Command{
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Use: "leoctl",
	}
)

func init() {
	pb.Wf = aw.New()
	_rootCmd.AddCommand(
		backend.Command(),
		ide.Command(),
		util.Command(),
		snippet.Command(),
	)
}

func main() {
	pb.Wf.Run(func() {
		_rootCmd.Execute()
		pb.Wf.SendFeedback()
	})
	// _rootCmd.Execute()
}
