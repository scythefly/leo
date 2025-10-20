package util

import (
	"leo/internal/pb"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var _randomOption struct {
	uuid   bool
	simple bool
}

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/!@#$%^&*()")
var schars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func randomCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "random",
		Run: func(_ *cobra.Command, args []string) {
			if _randomOption.uuid {
				u := uuid.NewString()
				pb.Wf.NewItem("> " + u).Copytext(u).Arg(u).Valid(true)
			} else {
				n := 8
				if len(args) > 0 {
					n, _ = strconv.Atoi(args[0])
					if n < 1 {
						n = 8
					}
					if n > 32 {
						n = 32
					}
				}
				p := chars
				if _randomOption.simple {
					p = schars
				}
				var out []byte
				for i := 0; i < n; i++ {
					out = append(out, p[rand.Int()%len(p)])
				}
				outString := string(out)
				pb.Wf.NewItem("> " + outString).Copytext(outString).Arg(outString).Valid(true)
			}
		},
	}

	flags := cmd.PersistentFlags()
	flags.BoolVar(&_randomOption.uuid, "uuid", false, "Use uuid")
	flags.BoolVar(&_randomOption.simple, "simple", false, "Use simple random string")

	return cmd
}
