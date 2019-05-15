package neocortex

import "context"

type CognitiveService interface {
	CreateNewContext(c *context.Context, info PersonInfo) *Context
	OnContextIsDone(callback func(c *Context))
	GetProtoResponse(c *Context, in *Input) (*Output, error)
}
