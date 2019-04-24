package neocortex

import (
	"time"
)

type TimelineOutputs map[time.Time]*Output
type TimelineInputs map[time.Time]*Input

type Dialog struct {
	ID      string          `json:"id" bson:"id"`
	Context *Context        `json:"context" bson:"context"`
	StartAt time.Time       `json:"start_at" bson:"start_at"`
	EndAt   time.Time       `json:"end_at" bson:"end_at"`
	Ins     TimelineInputs  `json:"ins" bson:"ins"`
	Outs    TimelineOutputs `json:"outs" bson:"outs"`
}
