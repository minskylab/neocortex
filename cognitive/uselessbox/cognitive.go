package uselessbox

import (
	"context"
	"github.com/bregydoc/neocortex"
)

type Cognitive struct {
	sessions map[string]string
}

func NewCognitive() *Cognitive {
	return &Cognitive{}
}

func (useless *Cognitive) CreateNewContext(c *context.Context, userID string) *neocortex.Context {
	return &neocortex.Context{
		Context:   c,
		SessionID: "useless session",
	}
}

func (useless *Cognitive) GetProtoResponse(c *neocortex.Context, in neocortex.Input) (neocortex.Output, error) {
	return NewOutput(c), nil
}
