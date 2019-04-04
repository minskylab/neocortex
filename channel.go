package neocortex

type MiddleHandler func(message Input, response OutputResponse) error
type ContextFabric func(userID string) *Context

type CommunicationChannel interface {
	RegisterMessageEndpoint(handler MiddleHandler) error
	ToHear() error
	GetContextFabric() ContextFabric // TODO: Rev
}
