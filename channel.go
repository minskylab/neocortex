package neocortex

import "context"

type CommunicationChannel interface {
	registerMessageEndpoint(func(c context.Context, message *Input) error) error
	sendResponse(c context.Context, message *Output) error
}
