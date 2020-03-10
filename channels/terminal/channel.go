package terminal

import (
	"bufio"
	"strconv"

	neo "github.com/minskylab/neocortex"
)

type Channel struct {
	reader               *bufio.Reader
	options              *ChannelOptions
	messageIn            neo.MiddleHandler      // req
	newContext           neo.ContextFabric      // req
	contexts             map[int64]*neo.Context // req
	newContextCallbacks  []*func(c *neo.Context)
	doneContextCallbacks []*func(c *neo.Context)
}

func (term *Channel) RegisterMessageEndpoint(handler neo.MiddleHandler) error {
	term.messageIn = handler
	return nil
}

func (term *Channel) ToHear() error {
	return term.renderUserInterface(false)
}

func (term *Channel) GetContextFabric() neo.ContextFabric {
	return term.newContext
}

func (term *Channel) SetContextFabric(fabric neo.ContextFabric) {
	term.newContext = fabric
}

func (term *Channel) OnNewContextCreated(callback func(c *neo.Context)) {
	if term.newContextCallbacks == nil {
		term.newContextCallbacks = []*func(c *neo.Context){}
	}
	term.newContextCallbacks = append(term.newContextCallbacks, &callback)
}

func (term *Channel) OnContextIsDone(callback func(c *neo.Context)) {
	if term.doneContextCallbacks == nil {
		term.doneContextCallbacks = []*func(c *neo.Context){}
	}
	term.doneContextCallbacks = append(term.doneContextCallbacks, &callback)
}

func (term *Channel) CallContextDone(c *neo.Context) {
	id, err := strconv.ParseInt(c.Person.ID, 10, 64)
	if err == nil {
		delete(term.contexts, id)
	}
}
