package analytics

import (
	"github.com/bregydoc/neocortex"
	"time"
)

type TimelineOutputs map[time.Time]neocortex.Output
type TimelineInputs map[time.Time]neocortex.Input

type Dialog struct {
	ID      string
	Context *neocortex.Context
	Ins     TimelineInputs
	Outs    TimelineOutputs
}
