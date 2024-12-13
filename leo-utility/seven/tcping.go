package seven

import (
    "fmt"
    "net"
    "time"

    aw "github.com/deanishe/awgo"
    "github.com/go-ping/ping"
)

// BuildTcpingFeedback ...
func BuildTcpingFeedback(wf *aw.Workflow, ipString, port string) {
    timeStart := time.Now()
    addr := ipString + ":"
    addr += port
    _, err := net.DialTimeout("tcp", addr, time.Second*3)
    response := time.Since(timeStart)
    if err != nil {
        wf.NewItem(fmt.Sprintf("> tcping, err: %s", err.Error()))
        return
    }
    wf.Rerun(1)
    wf.NewItem(fmt.Sprintf("> tcping, time=%v", response))
}

// BuildPingFeedback ...
func BuildPingFeedback(wf *aw.Workflow, ipString string) {
    pinger, err := ping.NewPinger(ipString)
    if err != nil {
        wf.NewItem(fmt.Sprintf("> ping, err, %s", err.Error()))
        return
    }
    wf.Rerun(1)
    pinger.Count = 1
    pinger.Run()
    stats := pinger.Statistics()
    if len(stats.Rtts) > 0 {
        wf.NewItem(fmt.Sprintf("> ping, time %v", stats.Rtts[0]))
        return
    }
    wf.NewItem("> ping, no result")
}
