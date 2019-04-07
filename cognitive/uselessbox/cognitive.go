package uselessbox

import (
	"context"
	neo "github.com/bregydoc/neocortex"
	"github.com/rs/xid"
)

type Cognitive struct{}

func NewCognitive() *Cognitive {
	return &Cognitive{}
}

func (useless *Cognitive) CreateNewContext(c *context.Context, info neo.PersonInfo) *neo.Context {
	id := xid.New()
	return &neo.Context{
		Context:   c,
		SessionID: id.String(),
	}
}

func (useless *Cognitive) GetProtoResponse(in *neo.Input) (*neo.Output, error) {
	if in.Context == nil {
		return nil, neo.ErrContextNotExist
	}
	return useless.NewOutputText(in.Context), nil
}
