package analytics

import (
	"github.com/bregydoc/neocortex"
	"time"
)

type TimelineOutputs map[time.Time]neocortex.Output
type TimelineInputs map[time.Time]neocortex.Input

type Dialog struct {
	ID      string             `json:"id"`
	Context *neocortex.Context `json:"context"`
	StartAt time.Time          `json:"start_at"`
	EndAt   time.Time          `json:"end_at"`
	Ins     TimelineInputs     `json:"ins"`
	Outs    TimelineOutputs    `json:"outs"`
}
