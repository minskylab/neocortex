package watson

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
	"github.com/watson-developer-cloud/go-sdk/core"
)

func (watson *Cognitive) NewInputText(c *neo.Context, text string, intents []neo.Intent, entities []neo.Entity) (*neo.Input, *assistantv2.MessageOptions) {
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
	options.SetContext(&assistantv2.MessageContext{
		Global: &assistantv2.MessageContextGlobal{
			System: &assistantv2.MessageContextGlobalSystem{
				UserID:   core.StringPtr(c.Person.ID),
				Timezone: core.StringPtr(c.Person.Timezone),
			},
		},
	})
	options.SetInput(input)

	data := neo.InputData{
		Type:  neo.InputText,
		Value: text,
		Data:  []byte(text),
	}

	return watson.NewInput(c, options, data), options
}
