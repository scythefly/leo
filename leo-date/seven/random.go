package seven

import (
	aw "github.com/deanishe/awgo"
	"github.com/google/uuid"
)

// BuildUUIDFeedback ...
func BuildUUIDFeedback(wf *aw.Workflow) {
	u, err := uuid.NewRandom()
	if err != nil {
		wf.NewItem(err.Error())
		return
	}
	wf.NewItem("> " + u.String()).Copytext(u.String()).Arg(u.String()).Valid(true)
}

// BuildRandomPswdFeedback ...
func BuildRandomPswdFeedback(wf *aw.Workflow, len int) {
}
