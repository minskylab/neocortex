package speech

import (
	"strconv"

	neo "github.com/minskylab/neocortex"
)

type Channel struct {
	options              *ChannelOptions
	messageIn            neo.MiddleHandler
	newContext           neo.ContextFabric
	contexts             map[int64]*neo.Context
	newContextCallbacks  []*func(c *neo.Context)
	doneContextCallbacks []*func(c *neo.Context)
}

type ChannelOptions struct {
	AccessToken string
	VerifyToken string
}

func (stts *Channel) RegisterMessageEndpoint(handler neo.MiddleHandler) error {
	stts.messageIn = handler
	return nil
}

func (stts *Channel) ToHear() error {
	// TODO
	return nil
}

func (stts *Channel) GetContextFabric() neo.ContextFabric {
	return stts.newContext
}

func (stts *Channel) SetContextFabric(fabric neo.ContextFabric) {
	stts.newContext = fabric
}

func (stts *Channel) OnNewContextCreated(callback func(c *neo.Context)) {
	if stts.newContextCallbacks == nil {
		stts.newContextCallbacks = []*func(c *neo.Context){}
	}
	stts.newContextCallbacks = append(stts.newContextCallbacks, &callback)
}

func (stts *Channel) OnContextIsDone(callback func(c *neo.Context)) {
	if stts.doneContextCallbacks == nil {
		stts.doneContextCallbacks = []*func(c *neo.Context){}
	}
	stts.doneContextCallbacks = append(stts.doneContextCallbacks, &callback)
}

func (stts *Channel) CallContextDone(c *neo.Context) {
	id, err := strconv.ParseInt(c.Person.ID, 10, 64)
	if err == nil {
		delete(stts.contexts, id)
	}
}
