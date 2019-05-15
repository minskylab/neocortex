package uselessbox

import (
	"context"

	neo "github.com/bregydoc/neocortex"
	"github.com/rs/xid"
)

type Cognitive struct {
	doneContextCallbacks []*func(c *neo.Context)
}

func NewCognitive() *Cognitive {
	return &Cognitive{}
}

func (useless *Cognitive) CreateNewContext(c *context.Context, info neo.PersonInfo) *neo.Context {
	id := xid.New()
	return &neo.Context{
		Context:   c,
		SessionID: id.String(),
		Person:    info,
		Variables: map[string]interface{}{},
	}
}

func (useless *Cognitive) GetProtoResponse(c *neo.Context, in *neo.Input) (*neo.Output, error) {
	if c == nil {
		return nil, neo.ErrContextNotExist
	}
	return useless.NewOutputText(c), nil
}

func (useless *Cognitive) OnContextIsDone(callback func(c *neo.Context)) {
	if useless.doneContextCallbacks == nil {
		useless.doneContextCallbacks = []*func(c *neo.Context){}
	}
	useless.doneContextCallbacks = append(useless.doneContextCallbacks, &callback)
}
