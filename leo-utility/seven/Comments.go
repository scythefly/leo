package seven

import aw "github.com/deanishe/awgo"

// BuildCommentsFeedback ...
func BuildCommentsFeedback(wf *aw.Workflow, c string) {
	var out string
	out = "-"
	out += c
	out += "-"
	for {
		if len(out) > 80 {
			break
		}
		out += " -"
		out += c
		out += "-"
	}
	wf.NewItem("> " + out).Arg(out).Copytext(out).Valid(true)
}
