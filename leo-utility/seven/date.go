package seven

import (
	"fmt"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/scythefly/orb/onet"
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

// BuildVpnExportFeedback ...
func BuildVpnExportFeedback(wf *aw.Workflow) {
	// export http_proxy=http://127.0.0.1:1087;export https_proxy=http://127.0.0.1:1087;
	ipString, err := onet.GetCurrentIP()
	if err != nil {
		wf.NewItem("> failed to get current ip")
		return
	}
	outString := fmt.Sprintf("export http_proxy=http://%s:1087;export https_proxy=http://%s:1087;", ipString, ipString)
	wf.NewItem("> " + outString).Copytext(outString).Arg(outString).Valid(true)
}
