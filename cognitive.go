package neocortex

import "context"

type CognitiveService interface {
	CreateNewContext(c *context.Context, info PersonInfo) *Context
	GetProtoResponse(in *Input) (*Output, error)
}
