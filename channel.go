package neocortex

type MiddleHandler func(message *Input, response func(output *Output) error) error

type CommunicationChannel interface {
	RegisterMessageEndpoint(handler MiddleHandler) error
}
