package facebook

import "github.com/bregydoc/neocortex"

type Input struct {
	c         *neocortex.Context
	inputType neocortex.InputType
	entities  []neocortex.Entity
	intents   []neocortex.Intent
}

func (i *Input) Context() *neocortex.Context {
	return i.c
}

func (i *Input) InputType() neocortex.InputType {
	return i.inputType
}

func (i *Input) Entities() []neocortex.Entity {
	return i.entities
}

func (i *Input) Intents() []neocortex.Intent {
	return i.intents
}
