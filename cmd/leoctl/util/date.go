package util

import (
	"fmt"
	"leo/internal/pb"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type option struct {
	unix bool
}

var _opt option

func dateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "date",
		Run: func(_ *cobra.Command, args []string) {
			t := time.Now()
			if len(args) > 0 && args[0] != "" {
				m := args[0][0]
				if m != '+' && m != '-' {
					timestamp, _ := strconv.Atoi(args[0])
					t = time.Unix(int64(timestamp), 0)
				} else {
					sub := 0
					duration := time.Second
					b := args[0][len(args[0])-1]
					switch b {
					case 'm':
						duration = time.Minute
					case 'h', 'H':
						duration = time.Hour
					case 'd', 'D':
						duration = 24 * time.Hour
					}
					if duration == time.Second {
						sub, _ = strconv.Atoi(args[0])
					} else {
						sub, _ = strconv.Atoi(args[0][0 : len(args[0])-1])
					}
					t = t.Add(time.Duration(sub) * duration)
				}
			}
			pb.Wf.Rerun(1)
			if _opt.unix {
				title := fmt.Sprintf("%d", t.Unix())
				pb.Wf.NewItem("> " + title).Subtitle(t.Format(time.DateTime)).Copytext(title).Arg(title).Valid(true)
			} else {
				title := t.Format(time.DateTime)
				pb.Wf.NewItem("> " + title).Copytext(title).Arg(title).Valid(true)
			}
		},
	}

	flags := cmd.PersistentFlags()
	flags.BoolVar(&_opt.unix, "unix", false, "Use unix timestamp")

	return cmd
}
