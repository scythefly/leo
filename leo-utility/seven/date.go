package seven

import (
	"fmt"
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/kakami/pkg/net"
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
func BuildVpnExportFeedback(wf *aw.Workflow, mark string) {
	// export http_proxy=http://127.0.0.1:1081;export https_proxy=http://127.0.0.1:1081;
	ips, err := net.CurrentIP()
	if err != nil {
		wf.NewItem("> failed to get current ip")
		return
	}
	var ipss []string
	ipss = append(ipss, "127.0.0.1")
	ipss = append(ipss, ips...)
	for idx := range ipss {
		if mark == "" || strings.Index(ipss[idx], mark) >= 0 {
			outString := fmt.Sprintf("export http_proxy=http://%s:1081;export https_proxy=http://%s:1081;", ipss[idx], ipss[idx])
			wf.NewItem("> " + outString).Copytext(outString).Arg(outString).Valid(true)
		}
	}
}
