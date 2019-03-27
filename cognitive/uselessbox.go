package cognitive

import (
	"context"
	"github.com/bregydoc/neocortex"
)

type UselessBoxCognitive struct {
	sessions map[string]string
}

func NewUselessBoxCognitive() *UselessBoxCognitive {
	return &UselessBoxCognitive{}
}

func (useless *UselessBoxCognitive) CreateNewSession(c context.Context, userID string) *neocortex.Context {
	return &neocortex.Context{
		Context:   c,
		SessionID: "useless session",
	}
}

func (useless *UselessBoxCognitive) GetProtoResponse(c *neocortex.Context, in *neocortex.Input) (*neocortex.Output, error) {
	return &neocortex.Output{
		Context: *c,
		Response: []*neocortex.ResponseGeneric{
			{ResponseType: neocortex.Text, Text: "I have problems, my brain is breaking"},
		},
		Intents: []*neocortex.Intent{},
	}, nil
}
