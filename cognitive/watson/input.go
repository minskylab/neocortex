package watson

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

func (watson *Cognitive) NewInput(c *neo.Context, opts *assistantv2.MessageOptions, inputType neo.InputType) *neo.Input {
	entities := make([]neo.Entity, 0)
	for _, e := range opts.Input.Entities {
		entities = append(entities, getNeocortexEntity(e))
	}

	intents := make([]neo.Intent, 0)
	for _, i := range opts.Input.Intents {
		intents = append(intents, getNeocortexIntent(i))
	}

	return &neo.Input{
		Context:   c,
		InputType: inputType,
		Intents:   intents,
		Entities:  entities,
	}
}
