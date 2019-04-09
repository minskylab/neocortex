package neocortex

import (
	"context"
	"github.com/sirupsen/logrus"
	"reflect"
)

func newDefaultEngine(cognitive CognitiveService, channels ...CommunicationChannel) *Engine {
	engine := &Engine{}
	engine.channels = channels
	engine.cognitive = cognitive
	engine.registeredResolvers = map[CommunicationChannel]map[Matcher]*HandleResolver{}
	engine.generalResolver = map[CommunicationChannel]*HandleResolver{}
	engine.done = make(chan error, 1)
	engine.logger = logrus.StandardLogger()
	engine.logger.SetLevel(logrus.DebugLevel)

	return engine
}

func New(cognitive CognitiveService, channels ...CommunicationChannel) (*Engine, error) {
	engine := newDefaultEngine(cognitive, channels...)
	fabric := func(ctx context.Context, info PersonInfo) *Context {
		return cognitive.CreateNewContext(&ctx, info)
	}

	for _, ch := range channels {
		engine.logger.Debug("Registering channel ", reflect.ValueOf(ch).Type())
		engine.registeredResolvers[ch] = map[Matcher]*HandleResolver{}

		ch.SetContextFabric(fabric)
		err := ch.RegisterMessageEndpoint(func(message *Input, response OutputResponse) error {
			return engine.onMessage(&ch, message, response)
		})

		if err != nil {
			return nil, err
		}

		go func(ch CommunicationChannel) {
			err := ch.ToHear()
			if err != nil {
				engine.done <- err
			}
		}(ch)
	}

	return engine, nil
}

func (engine *Engine) Run() error {
	return <-engine.done
}
