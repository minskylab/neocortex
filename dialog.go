package neocortex

import (
	"time"
)

type TimelineOutputs map[time.Time]Output
type TimelineInputs map[time.Time]Input
type TimelineContexts map[time.Time]Context

type Dialog struct {
	ID           string           `json:"id"`
	StartAt      time.Time        `json:"start_at"`
	LastActivity time.Time        `json:"last_activity"`
	EndAt        time.Time        `json:"end_at"`
	Ins          TimelineInputs   `json:"ins"`
	Outs         TimelineOutputs  `json:"outs"`
	Contexts     TimelineContexts `json:"contexts"`
}
