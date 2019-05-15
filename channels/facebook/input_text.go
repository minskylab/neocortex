package facebook

import neo "github.com/bregydoc/neocortex"

func (fb *Channel) NewInputText(text string, i []neo.Intent, e []neo.Entity) *neo.Input {
	t := neo.InputData{
		Type:  neo.InputText,
		Value: text,
		Data:  []byte(text),
	}
	return fb.NewInput(t, i, e)
}
