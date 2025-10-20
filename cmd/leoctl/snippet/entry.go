package snippet

import (
	"fmt"
	"leo/internal/biz/snippet"
	"leo/internal/pb"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "snippet",
		Run: query,
	}
	cmd.AddCommand(
		addCommand(),
		rmCommand(),
	)

	return cmd
}

func query(_ *cobra.Command, args []string) {
	if len(args) < 1 {
		pb.Wf.NewItem("query with keyword, or add/rm").Valid(false)
		return
	}
	var key, value string
	key = args[0]
	if len(args) > 1 {
		value = args[1]
	}
	pairs := snippet.Query(key, value)
	if len(pairs) == 0 {
		pb.Wf.NewItem("no snippet found").Valid(false)
		return
	}

	for _, pair := range pairs {
		pb.Wf.NewItem("> " + pair.Value).
			Subtitle(fmt.Sprintf("%s -> %s", pair.Key, pair.Value)).
			Copytext(pair.Value).Arg(pair.Value).Valid(true)
	}
}

func addCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add",
		Run: add,
	}

	return cmd
}

func add(_ *cobra.Command, args []string) {
	it := pb.Wf.NewItem("Add ")
	if len(args) < 1 {
		it.Subtitle("insert value to db with key 'Shift'").Valid(false)
		return
	}

	key := args[0]
	value := key
	if len(args) > 1 {
		value = args[1]
	}

	it.Arg(fmt.Sprintf("snippet add %s %s", key, value))
	it.NewModifier("shift").Subtitle(fmt.Sprintf("add %s -> %s", key, value)).Valid(true)
}

func rmCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "rm",
		Run: rm,
	}

	return cmd
}

func rm(_ *cobra.Command, args []string) {
	if len(args) < 1 {
		it := pb.Wf.NewItem("Remove ")
		it.Subtitle("remove from db with key 'Shift'").Valid(false)
		return
	}
	var value string
	if len(args) > 1 {
		value = args[1]
	}

	pairs := snippet.Query(args[0], value)

	for _, pair := range pairs {
		it := pb.Wf.NewItem("Remove")
		it.Arg(fmt.Sprintf("snippet rm %s %s", pair.Key, pair.Value)).
			Subtitle(fmt.Sprintf("remove %s -> %s", pair.Key, pair.Value)).Valid(true)
		it.NewModifier("shift").Subtitle(fmt.Sprintf("do remove %s -> %s", pair.Key, pair.Value)).Valid(true)
	}
}
