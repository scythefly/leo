package main

import (
	"strconv"

	aw "github.com/deanishe/awgo"

	"leo-utility/seven"
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
		if intervals > 32 {
			intervals = 32
		}
		seven.BuildRandomPswdFeedback(wf, intervals)
	}

	// ip query
	if query == "ip" || query == "ip " {
		if argsLen > 1 {
			seven.BuildIPQueryFeedback(wf, args[1])
		}
	}

	// vpn export
	if query == "vpn" || query == "vpn " {
		seven.BuildVpnExportFeedback(wf)
	}

	// ping
	if query == "ping" || query == "ping " {
		if argsLen > 2 {
			seven.BuildTcpingFeedback(wf, args[1], args[2])
		} else if argsLen > 1 {
			seven.BuildPingFeedback(wf, args[1])
		}
	}

	if query == "comments" || query == "comments " {
		if argsLen > 1 {
			seven.BuildCommentsFeedback(wf, args[1])
		} else {
			seven.BuildCommentsFeedback(wf, "-")
		}
	}

	if query == "ide" || query == "ide " {
		if argsLen > 2 {
			seven.BuildIDEFeedback(wf, args[1], args[2])
		} else if argsLen > 1 {
			seven.BuildIDEFeedback(wf, args[1], "-")
		}
	}

	if wf.IsEmpty() {
		wf.NewItem("ld/lt +/- intervals")
		wf.NewItem("luu -> uuid")
		wf.NewItem("lp length -> password")
		wf.NewItem("ip xxx.xxx.xxx.xxx")
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
