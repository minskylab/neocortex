package neocortex

import "context"

// Context represent the context of a conversation
type Context struct {
	SessionID string
	Context   context.Context
	Metadata  map[string]string
}
