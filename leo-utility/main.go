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
		if argsLen > 1 {
			seven.BuildVpnExportFeedback(wf, args[1])
		} else {
			seven.BuildVpnExportFeedback(wf, "")
		}
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

	if query == "command" || query == "command " {
		if argsLen > 2 {
			p := args[2:]
			seven.BuildCommandFeedback(wf, args[1], p...)
		} else {
			wf.NewItem("go on inputing...")
		}
	}

	if query == "proxy" || query == "proxy " {
		if argsLen > 1 {
			if args[1] == "start" || args[1] == "stop" {
				wf.NewItem("> local trojan " + args[1]).Arg(args[1]).Valid(true)
			} else if args[1] == "global" ||
				args[1] == "pac" || args[1] == "manual" {
				wf.NewItem("> set proxy model to " + args[1]).Arg(args[1]).Valid(true)
			} else {
				wf.NewItem("start - stop - global - pac - manual")
			}
		} else {
			wf.NewItem("start - stop - global - pac - manual")
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
