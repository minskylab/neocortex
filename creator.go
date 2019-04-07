package neocortex

import (
	"context"
	"github.com/sirupsen/logrus"
)

func New(cognitive CognitiveService, channels ...CommunicationChannel) (*Engine, error) {
	engine := &Engine{}
	engine.channels = channels
	engine.cognitive = cognitive
	engine.registeredResolvers = map[CommunicationChannel]map[Matcher]*HandleResolver{}
	engine.done = make(chan error, 1)
	engine.logger = logrus.StandardLogger()
	engine.logger.SetLevel(logrus.DebugLevel)

	fabric := func(ctx context.Context, info PersonInfo) *Context {
		return cognitive.CreateNewContext(&ctx, info)
	}

	for _, ch := range channels {
		ch.SetContextFabric(fabric)
		err := ch.RegisterMessageEndpoint(func(message *Input, response OutputResponse) error {
			return engine.onMessage(ch, message, response)
		})

		if err != nil {
			return nil, err
		}

		go func() {
			err := ch.ToHear()
			if err != nil {
				engine.done <- err
			}
		}()
	}

	return engine, nil
}

func (engine *Engine) Run() error {
	return <-engine.done
}
