package uselessbox

import neo "github.com/bregydoc/neocortex"

func (useless *Cognitive) NewInput(data neo.InputData, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Data:     data,
		Intents:  i,
		Entities: e,
	}
}
