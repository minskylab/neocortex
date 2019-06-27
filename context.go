package neocortex

import (
	"context"
	"fmt"
)

// PersonInfo describes the basic info of the person
type PersonInfo struct {
	ID       string `json:"id"`
	Timezone string `json:"timezone"`
	Picture  string `json:"picture"`
	Locale   string `json:"locale"`
	Name     string `json:"name"`
}

// Context represent the context of a conversation
type Context struct {
	Context   *context.Context       `json:"-" bson:"-"`
	SessionID string                 `json:"session_id" bson:"session_id"`
	Person    PersonInfo             `json:"person" bson:"person"`
	Variables map[string]interface{} `json:"variables" bson:"variables"`
}

func (c *Context) String() string {
	s := "\n===== CONTEXT =====\n"
	s = s + fmt.Sprintf("session: %s\n", c.SessionID)
	s = s + fmt.Sprintf("context: %v\n", c.Context)
	s = s + fmt.Sprintf("user name: %s\n", c.Person.Name)
	s = s + fmt.Sprintf("total context variables: %d\n", len(c.Variables))
	s = s + "======================\n"
	return s
}
