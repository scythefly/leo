package main

import (
	"leo/cmd/leoctl/backend"
	"leo/cmd/leoctl/ide"
	"leo/cmd/leoctl/snippet"

	aw "github.com/deanishe/awgo"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

// func main() {
// 	dir, _ := os.UserHomeDir()
// 	socketPath := dir + "/.peon/leo.sock"
// 	conn, err := grpc.NewClient("unix://"+socketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		panic(err)
// 	}
// 	client := snippet.NewSnippetClient(conn)
// 	resp, _ := client.Query(context.Background(), &snippet.Request{Key: "22"})
// 	fmt.Println(resp)
// }

func main() {
	wf.Run(run)
}

func run() {
	args := wf.Args()
	if len(args) == 0 {
		wf.NewItem("select a tool...").Valid(false)
		wf.SendFeedback()
		return
	}

	switch args[0] {
	case "snippet":
		snippet.BuildFeedback(wf)
	case "backend":
		backend.BuildFeedback(wf)
	case "ide":
		ide.BuildFeedback(wf)
	default:
		str := args[0]
		for i := 1; i < len(args); i++ {
			str += "|" + args[i]
		}
		wf.NewItem("unknown command: " + str).Valid(false)
	}

	wf.SendFeedback()
}
