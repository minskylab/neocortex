package neocortex

import (
	"context"
)

func New(cognitive CognitiveService, channels ...CommunicationChannel) (*Engine, error) {
	middle := &Engine{}
	middle.channels = channels
	middle.cognitive = cognitive
	middle.registeredResolvers = map[CommunicationChannel]map[Matcher]*HandleResolver{}
	middle.done = make(chan error, 1)
	fabric := func(ctx context.Context, userID string) *Context {
		return cognitive.CreateNewContext(&ctx, userID)
	}

	for _, ch := range channels {
		ch.SetContextFabric(fabric)
		err := ch.RegisterMessageEndpoint(func(message *Input, response OutputResponse) error {
			return middle.onMessage(ch, message, response)
		})

		if err != nil {
			return nil, err
		}

		go func() {
			err := ch.ToHear()
			if err != nil {
				middle.done <- err
			}
		}()
	}

	return middle, nil
}

func (cortex *Engine) Run() error {
	return <-cortex.done
}
