package seven

import (
	"fmt"
	"time"

	aw "github.com/deanishe/awgo"
)

// BuildDateFeedback ...
func BuildDateFeedback(wf *aw.Workflow, intervals int) {
	t := int(time.Now().Unix())
	to := t + intervals
	title := fmt.Sprintf("%d", to)
	wf.NewItem("> " + title).Arg(title).Copytext(title).Valid(true)
}

// BuildTimeStringFeedback ...
func BuildTimeStringFeedback(wf *aw.Workflow, tm int) {
	if tm == 0 {
		wf.Rerun(1)
		t := time.Now()
		title := t.Format("2006-01-02 15:04:05")
		wf.NewItem("> " + title).Copytext(title).Arg(title).Valid(true)
	} else {
		t := time.Unix(int64(tm), 0)
		title := t.Format("2006-01-02 15:04:05")
		wf.NewItem("> " + title).Copytext(title).Arg(title).Valid(true)
	}
}
