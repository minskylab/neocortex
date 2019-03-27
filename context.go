package neocortex

import "context"

// Context represent the context of a conversation
type Context struct {
	Context   *context.Context
	SessionID string
	Variables map[string]interface{}
}
