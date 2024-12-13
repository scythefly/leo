package seven

import aw "github.com/deanishe/awgo"

// BuildCommentsFeedback ...
func BuildCommentsFeedback(wf *aw.Workflow, c string) {
	var out string
	prefix := (78 - len(c)) / 2
	for i := 0; i < prefix; i++ {
		out += "-"
	}
	out += " "
	out += c
	out += " "
	for {
		if len(out) > 80 {
			break
		}
		out += "-"
	}
	wf.NewItem("> " + out).Arg(out).Copytext(out).Valid(true)
}
