package neocortex

import (
	"context"
)

type PersonInfo struct {
	ID       string `json:"id"`
	Timezone string `json:"timezone"`
	Name     string `json:"name"`
}

// Context represent the context of a conversation
type Context struct {
	Context   *context.Context       `json:"-" bson:"-"`
	SessionID string                 `json:"session_id" bson:"session_id"`
	Person    PersonInfo             `json:"person" bson:"person"`
	Variables map[string]interface{} `json:"variables" bson:"variables"`
}
