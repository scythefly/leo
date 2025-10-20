package date

import aw "github.com/deanishe/awgo"

func BuildFeedback(wf *aw.Workflow) {
	// args[1] path
	// args[2] project
	args := wf.Args()
	if len(args) < 3 || args[2] == "" {
		wf.NewItem("code $project").Valid(false)
		return
	}
}
