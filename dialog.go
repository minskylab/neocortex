package neocortex

import (
	"time"
)

type TimelineOutputs map[time.Time]Output
type TimelineInputs map[time.Time]Input

type Dialog struct {
	ID      string          `json:"id"`
	Context *Context        `json:"context"`
	StartAt time.Time       `json:"start_at"`
	EndAt   time.Time       `json:"end_at"`
	Ins     TimelineInputs  `json:"ins"`
	Outs    TimelineOutputs `json:"outs"`
}
