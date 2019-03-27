package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

type Input struct {
	context   *neocortex.Context
	opts      *assistantv2.MessageOptions
	inputType neocortex.InputType
}

func (i *Input) Context() *neocortex.Context {
	i.context.SessionID = *i.opts.SessionID
	return i.context
}

func (i *Input) InputType() neocortex.InputType {
	return i.inputType
}

func (i *Input) Entities() []neocortex.Entity {
	entities := make([]neocortex.Entity, 0)
	for _, e := range i.opts.Input.Entities {
		entities = append(entities, getNeocortexEntity(e))
	}
	return entities
}

func (i *Input) Intents() []neocortex.Intent {
	intents := make([]neocortex.Intent, 0)
	for _, intent := range i.opts.Input.Intents {
		intents = append(intents, getNeocortexIntent(intent))
	}
	return intents
}
