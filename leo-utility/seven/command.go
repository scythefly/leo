package seven

import (
	"fmt"

	aw "github.com/deanishe/awgo"
)

func BuildCommandFeedback(wf *aw.Workflow, cmd string, args ...string) {
	var outString = "go on inputing..."
	switch cmd {
	case "ps":
		outString = buildPSString(args...)
	case "pk":
		outString = buildPSString(args...)
		outString = fmt.Sprintf("kill -9 `%s | awk '{print $2}'`", outString)
	}
	wf.NewItem("> " + outString).Copytext(outString).Arg(outString).Valid(true)
}

func buildPSString(args ...string) string {
	out := "ps aux "
	for _, arg := range args {
		out += fmt.Sprintf("| grep \"%s\" ", arg)
	}
	out += " | grep -v \"grep\""
	return out
}
