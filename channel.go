package neocortex

type CommunicationChannel interface {
	RegisterMessageEndpoint(handler MiddleHandler) error
	ToHear() error
	GetContextFabric() ContextFabric // TODO: Rev
	SetContextFabric(fabric ContextFabric)
	OnNewContextCreated(callback func(c *Context))
	OnContextIsDone(callback func(c *Context))
	CallContextDone(c *Context)
}
