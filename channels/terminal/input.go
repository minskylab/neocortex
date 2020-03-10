package terminal

import neo "github.com/minskylab/neocortex"

func (term *Channel) NewInput(data neo.InputData, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Data:     data,
		Intents:  i,
		Entities: e,
	}
}
