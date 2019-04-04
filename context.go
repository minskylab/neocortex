package neocortex

import "context"

// Context represent the context of a conversation
type Context struct {
	Context   *context.Context       `json:"-"`
	SessionID string                 `json:"session_id"`
	UserID    string                 `json:"user_id"`
	Variables map[string]interface{} `json:"variables"`
}
