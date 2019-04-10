package neocortex

import (
	"context"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
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

	cognitive.OnContextIsDone(func(c *Context) {
		engine.ActiveDialogs[c].EndAt = time.Now()
		_, err := engine.Repository.SaveNewDialog(engine.ActiveDialogs[c])
		delete(engine.ActiveDialogs, c)
		if err != nil {
			engine.done <- err
		}
	})

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

		ch.OnNewContextCreated(func(c *Context) {
			engine.ActiveDialogs[c] = &Dialog{
				ID:      xid.New().String(),
				Context: c,
				StartAt: time.Now(),
				EndAt:   nil,
				Ins:     TimelineInputs{},
				Outs:    TimelineOutputs{},
			}
		})

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
