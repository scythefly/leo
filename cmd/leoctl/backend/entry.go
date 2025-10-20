package backend

import (
	"fmt"
	"leo/cmd/leoctl/snippet"

	aw "github.com/deanishe/awgo"
)

func BuildFeedback(wf *aw.Workflow) {
	// args[0] == "backend"
	args := wf.Args()
	str := args[0]
	for i := 1; i < len(args); i++ {
		str += "|" + args[i]
	}
	fmt.Printf("backend command: %s\n", str)

	if len(args) > 4 {
		switch args[1] {
		case "snippet":
			switch args[2] {
			case "add":
				snippet.Add(args[3], args[4])
			case "rm":
				snippet.Rm(args[3], args[4])
			}
		}
	}
}
