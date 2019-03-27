package neocortex

type MiddleHandler func(message *Input, response OutputResponse) error

type CommunicationChannel interface {
	RegisterMessageEndpoint(handler MiddleHandler) error
	LaunchAndWait() error
}
