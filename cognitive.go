package neocortex

import "context"

type CognitiveService interface {
	CreateNewContext(c *context.Context, userID string) *Context
	GetProtoResponse(in *Input) (*Output, error)
}
