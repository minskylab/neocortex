package neocortex

import (
	"time"

	"github.com/rs/xid"
)

type InputRecord struct {
	At    time.Time `json:"at"`
	Input Input     `json:"input"`
}

type OutputRecord struct {
	At     time.Time `json:"at"`
	Output Output    `json:"output"`
}

type ContextRecord struct {
	At      time.Time `json:"at"`
	Context Context   `json:"context"`
}

type Dialog struct {
	ID           string           `json:"id" bson:"id"`
	StartAt      time.Time        `json:"start_at" bson:"start_at"`
	LastActivity time.Time        `json:"last_activity" bson:"last_activity"`
	EndAt        time.Time        `json:"end_at" bson:"end_at"`
	Ins          []*InputRecord   `json:"ins" bson:"ins"`
	Outs         []*OutputRecord  `json:"outs" bson:"outs"`
	Contexts     []*ContextRecord `json:"contexts" bson:"contexts"`
	Performance  float64          `json:"performance" bson:"performance"`
}

func newDialog() *Dialog {
	return &Dialog{
		ID:           xid.New().String(),
		LastActivity: time.Now(),
		StartAt:      time.Now(),
		EndAt:        time.Time{},
		Ins:          []*InputRecord{},
		Outs:         []*OutputRecord{},
		Contexts:     []*ContextRecord{},
		Performance:  0.0,
	}
}

// TODO: To optimize, please find all at the same time (that has better performance)
func (dialog *Dialog) HasEntity(entity string) bool {
	totalIns := len(dialog.Ins)
	totalOuts := len(dialog.Outs)

	diff := totalIns - totalOuts

	if diff == 0 {
		for i := 0; i < totalIns; i++ {
			for _, ent := range dialog.Ins[i].Input.Entities {
				if ent.Value == entity {
					return true
				}
			}
			for _, ent := range dialog.Outs[i].Output.Entities {
				if ent.Value == entity {
					return true
				}
			}
		}
		return false
	}

	for _, in := range dialog.Ins {
		for _, ent := range in.Input.Entities {
			if ent.Value == entity {
				return true
			}
		}
	}

	for _, out := range dialog.Outs {
		for _, ent := range out.Output.Entities {
			if ent.Value == entity {
				return true
			}
		}
	}

	return false
}

func (dialog *Dialog) HasIntent(intent string) bool {
	totalIns := len(dialog.Ins)
	totalOuts := len(dialog.Outs)

	diff := totalIns - totalOuts

	if diff == 0 {
		for i := 0; i < totalIns; i++ {
			for _, intn := range dialog.Ins[i].Input.Intents {
				if intn.Intent == intent {
					return true
				}
			}
			for _, intn := range dialog.Outs[i].Output.Intents {
				if intn.Intent == intent {
					return true
				}
			}
		}
		return false
	}

	for _, in := range dialog.Ins {
		for _, intn := range in.Input.Intents {
			if intn.Intent == intent {
				return true
			}
		}
	}

	for _, out := range dialog.Outs {
		for _, intn := range out.Output.Intents {
			if intn.Intent == intent {
				return true
			}
		}
	}

	return false
}

func (dialog *Dialog) HasDialogNode(node string) bool {
	for _, out := range dialog.Outs {
		for _, visited := range out.Output.VisitedNodes {
			if visited.Name == node || visited.Title == node {
				return true
			}
		}
	}

	return false
}
