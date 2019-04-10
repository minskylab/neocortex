package neocortex

import (
	"context"
	"os"
	"os/signal"

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
	// engine.logger.SetLevel(logrus.DebugLevel)
	engine.ActiveDialogs = map[*Context]*Dialog{}
	return engine
}

func New(repository Repository, cognitive CognitiveService, channels ...CommunicationChannel) (*Engine, error) {
	engine := newDefaultEngine(cognitive, channels...)
	engine.Repository = repository
	fabric := func(ctx context.Context, info PersonInfo) *Context {
		return cognitive.CreateNewContext(&ctx, info)
	}

	cognitive.OnContextIsDone(func(c *Context) {
		engine.OnContextIsDone(c)
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
			engine.OnNewContextCreated(c)
		})

		ch.OnContextIsDone(func(c *Context) {
			engine.OnContextIsDone(c)
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
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		engine.logger.Infoln("Closing all dialogs")
		for _, d := range engine.ActiveDialogs {
			engine.OnContextIsDone(d.Context)
		}
		engine.done <- nil
	}()
	return <-engine.done
}
