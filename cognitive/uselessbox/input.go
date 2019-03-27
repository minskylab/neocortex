package uselessbox

import "github.com/bregydoc/neocortex"

type Input struct {
	context   *neocortex.Context
	inputType neocortex.InputType
}

func (i *Input) Context() *neocortex.Context {
	return i.context
}

func (i *Input) InputType() neocortex.InputType {
	return i.inputType
}

func (i *Input) Entities() []neocortex.Entity {
	return []neocortex.Entity{}
}

func (i *Input) Intents() []neocortex.Intent {
	return []neocortex.Intent{}
}
