package telegram

import (
	"github.com/minskylab/neocortex"
)

type Channel struct {}

func (c *Channel) RegisterMessageEndpoint(handler neocortex.MiddleHandler) error {
	panic("implement me")
}

func (c *Channel) ToHear() error {
	panic("implement me")
}

func (c *Channel) GetContextFabric() neocortex.ContextFabric {
	panic("implement me")
}

func (c *Channel) SetContextFabric(fabric neocortex.ContextFabric) {
	panic("implement me")
}

func (c *Channel) OnNewContextCreated(callback func(c *neocortex.Context)) {
	panic("implement me")
}

func (c *Channel) OnContextIsDone(callback func(c *neocortex.Context)) {
	panic("implement me")
}

func (c *Channel) CallContextDone(c *neocortex.Context) {
	panic("implement me")
}
