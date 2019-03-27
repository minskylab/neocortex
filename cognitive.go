package neocortex

import "context"

type CognitiveService interface {
	CreateNewSession(c context.Context, userID string) *Context
	GetProtoResponse(c *Context, in *Input) (*Output, error)
}
