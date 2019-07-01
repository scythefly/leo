package main

import (
	"strconv"

	aw "github.com/deanishe/awgo"
	"github.com/scythefly/leo/leo-date/seven"
)

var (
	helpURL = "https://www.google.com"
	wf      *aw.Workflow
)

func init() {
	wf = aw.New(aw.HelpURL(helpURL))
}

func run() {
	var query string
	var intervals int
	args := wf.Args()
	argsLen := len(args)

	if argsLen > 0 {
		query = args[0]
	}

	if query == "date" || query == "date " {
		if argsLen > 1 {
			intervals, _ = strconv.Atoi(args[1])
		}
		wf.Rerun(1)
		seven.BuildDateFeedback(wf, intervals)
	}

	if query == "unix" || query == "unix " {
		if argsLen > 1 {
			intervals, _ = strconv.Atoi(args[1])
		}
		if intervals < 0 {
			wf.NewItem("timestamp < 0, error")
		} else {
			seven.BuildTimeStringFeedback(wf, intervals)
		}
	}

	if query == "uuid" || query == "uuid " {
		seven.BuildUUIDFeedback(wf)
	}

	if query == "password" || query == "password " {
		if argsLen > 1 {
			intervals, _ = strconv.Atoi(args[1])
		}
		if intervals < 1 {
			intervals = 8
		}
		seven.BuildRandomPswdFeedback(wf, intervals)
	}

	if wf.IsEmpty() {
		wf.NewItem("ld/lt +/- intervals")
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
