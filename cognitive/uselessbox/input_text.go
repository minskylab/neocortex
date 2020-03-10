package uselessbox

import neo "github.com/minskylab/neocortex"

func (useless *Cognitive) NewInputText(text string, i []neo.Intent, e []neo.Entity) *neo.Input {
	t := neo.InputData{
		Type:  neo.InputText,
		Value: text,
		Data:  []byte(text),
	}
	return useless.NewInput(t, i, e)
}
