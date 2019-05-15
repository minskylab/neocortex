package watson

import (
	neo "github.com/bregydoc/neocortex"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
	"github.com/watson-developer-cloud/go-sdk/core"
)

func (watson *Cognitive) NewInputText(text string, c *neo.Context, intents []neo.Intent, entities []neo.Entity) (*neo.Input, *assistantv2.MessageOptions) {
	ets := make([]assistantv2.RuntimeEntity, 0)
	for _, e := range entities {
		ets = append(ets, getNativeEntity(&e))
	}

	its := make([]assistantv2.RuntimeIntent, 0)
	for _, i := range intents {
		its = append(its, getNativeIntent(&i))
	}

	options := &assistantv2.MessageOptions{
		AssistantID: core.StringPtr(watson.assistantID),
		SessionID:   core.StringPtr(c.SessionID),
		Context: &assistantv2.MessageContext{
			Global: &assistantv2.MessageContextGlobal{
				System: &assistantv2.MessageContextGlobalSystem{
					UserID:   core.StringPtr(c.Person.ID),
					Timezone: core.StringPtr(c.Person.Timezone),
					// TurnCount: core.Int64Ptr(int64(watson.turnsMap[c.SessionID])),
				},
			},
			Skills: &assistantv2.MessageContextSkills{
				"main skill": map[string]interface{}{
					"user_defined": c.Variables,
				},
			},
		},
		Input: &assistantv2.MessageInput{
			MessageType: core.StringPtr("text"),
			Text:        &text,
			Intents:     its,
			Entities:    ets,
			Options: &assistantv2.MessageInputOptions{

				Debug:         core.BoolPtr(true),
				ReturnContext: core.BoolPtr(true),
			},
		},
	}

	data := neo.InputData{
		Type:  neo.InputText,
		Value: text,
		Data:  []byte(text),
	}

	return watson.NewInput(options, data), options
}
