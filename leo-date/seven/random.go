package seven

import (
	"math/rand"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/google/uuid"
)

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/!@#$%^&*()")

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
func BuildRandomPswdFeedback(wf *aw.Workflow, length int) {
	var out []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		out = append(out, chars[rand.Int()%len(chars)])
	}
	outString := string(out)
	wf.NewItem("> " + outString).Copytext(outString).Arg(outString).Valid(true)
}
