package snippet

import (
	"context"
	"fmt"
	"leo/api/snippet/v1"
	"os"

	aw "github.com/deanishe/awgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BuildFeedback(wf *aw.Workflow) {
	args := wf.Args()
	if len(args) < 2 {
		wf.NewItem("please input a keyword, or add/rm").Valid(false)
		return
	}

	switch args[1] {
	case "add":
		it := wf.NewItem("Add ")
		it.Subtitle("insert value to db with key 'Shift'").Valid(false)
		if len(args) > 3 {
			it.Arg(fmt.Sprintf("snippet add %s %s", args[2], args[3]))
			it.NewModifier("shift").Subtitle(fmt.Sprintf("add %s -> %s", args[2], args[3])).Valid(true)
		} else if len(args) > 2 {
			it.Arg(fmt.Sprintf("snippet add %s ", args[2]))
			it.NewModifier("shift").Subtitle(fmt.Sprintf("add %s -> %s", args[2], args[2])).Valid(true)
		}
	case "rm":
		if len(args) > 2 {
			buildRm(wf, args[2])
		} else {
			wf.NewItem("Remove").Subtitle("remove from db with key 'Shift'").Valid(false)
		}
	}

	req := &snippet.Request{
		Key: args[1],
	}
	if len(args) > 2 {
		req.Value = args[2]
	}

	resp, err := client().Query(context.Background(), req)
	if err != nil {
		wf.NewItem("error: " + err.Error()).Valid(false)
		return
	}

	for _, pair := range resp.GetPairs() {
		wf.NewItem("> " + pair.Value).Subtitle(fmt.Sprintf("%s -> %s", pair.Key, pair.Value)).
			Copytext(pair.Value).Arg(pair.Value).Valid(true)
	}
}

func buildRm(wf *aw.Workflow, key string) {
	pattern := "%" + key + "%"
	req := &snippet.Request{
		Key: pattern,
	}
	resp, err := client().Query(context.Background(), req)
	if err != nil {
		wf.NewItem("error: " + err.Error()).Valid(false)
		return
	}

	for _, pair := range resp.GetPairs() {
		it := wf.NewItem("Remove")
		it.Arg(fmt.Sprintf("snippet rm %s %s", pair.Key, pair.Value)).
			Subtitle(fmt.Sprintf("remove %s -> %s", pair.Key, pair.Value)).Valid(true)
		it.NewModifier("shift").Subtitle(fmt.Sprintf("do remove %s -> %s", pair.Key, pair.Value))
	}
}

func Add(key, value string) {
	client().Put(context.Background(), &snippet.Request{
		Key:   key,
		Value: value,
	})
}

func Rm(key, value string) {
	client().Delete(context.Background(), &snippet.Request{
		Key:   key,
		Value: value,
	})
}

func client() snippet.SnippetClient {
	dir, _ := os.UserHomeDir()
	socketPath := dir + "/.peon/leo.sock"
	conn, err := grpc.NewClient("unix://"+socketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return snippet.NewSnippetClient(conn)
}
