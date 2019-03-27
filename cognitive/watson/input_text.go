package watson

import (
	"github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
	"github.com/watson-developer-cloud/go-sdk/core"
)

type InputText struct {
	opts *assistantv2.MessageOptions
}

func (i *InputText) Type() neocortex.PrimitiveInputType {
	return neocortex.PrimitiveInputText
}

func (i *InputText) Value() string {
	return *i.opts.Input.Text
}

func (i *InputText) Data() []byte {
	return []byte(*i.opts.Input.Text)
}

func (watson *Cognitive) NewInputText(c *neocortex.Context, text string, intents []neocortex.Intent, entities []neocortex.Entity) *Input {
	ets := make([]assistantv2.RuntimeEntity, 0)
	for _, e := range entities {
		ets = append(ets, getNativeEntity(&e))
	}

	its := make([]assistantv2.RuntimeIntent, 0)
	for _, i := range intents {
		its = append(its, getNativeIntent(&i))
	}

	input := &assistantv2.MessageInput{
		MessageType: core.StringPtr("text"),
		Text:        &text,
		Intents:     its,
		Entities:    ets,
		Options: &assistantv2.MessageInputOptions{
			Debug:         core.BoolPtr(true),
			ReturnContext: core.BoolPtr(true),
		},
	}

	options := watson.service.NewMessageOptions(watson.assistantID, c.SessionID)

	options.SetInput(input)

	return &Input{
		context:   c,
		opts:      options,
		inputType: &InputText{options},
	}
}
