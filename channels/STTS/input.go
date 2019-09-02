package speech

import (
	neo "github.com/bregydoc/neocortex"
)

func (stts *Channel) NewInput(data neo.InputData, i []neo.Intent, e []neo.Entity) *neo.Input {
	return &neo.Input{
		Data:     data,
		Intents:  i,
		Entities: e,
	}
}
