package neocortex

import (
	"context"

	"os"
	"os/signal"
)

func newDefaultEngine(cognitive CognitiveService, channels ...CommunicationChannel) *Engine {
	engine := &Engine{}
	engine.channels = channels
	engine.cognitive = cognitive
	engine.registeredResolvers = map[CommunicationChannel]map[*Matcher]*HandleResolver{}
	engine.generalResolver = map[CommunicationChannel]*HandleResolver{}
	engine.registeredInjection = map[CommunicationChannel]map[*Matcher]*InInjection{}
	engine.generalInjection = map[CommunicationChannel]*InInjection{}
	engine.done = make(chan error, 1)
	// engine.logger = logrus.StandardLogger() // In the future
	engine.ActiveDialogs = map[*Context]*Dialog{}
	return engine
}

func Default(repository Repository, cognitive CognitiveService, channels ...CommunicationChannel) (*Engine, error) {
	engine := newDefaultEngine(cognitive, channels...)
	engine.Repository = repository
	engine.api = API{
		repository: repository,
		prefix:     "/api",
		Port:       ":4200",
	}

	fabric := func(ctx context.Context, info PersonInfo) *Context {
		return cognitive.CreateNewContext(&ctx, info)
	}

	cognitive.OnContextIsDone(func(c *Context) {
		engine.OnContextIsDone(c)
	})

	for _, ch := range channels {
		// engine.logger.Debug("Registering channel ", reflect.ValueOf(ch).Type())
		engine.registeredResolvers[ch] = map[*Matcher]*HandleResolver{}

		ch.SetContextFabric(fabric)
		err := ch.RegisterMessageEndpoint(func(message *Input, response OutputResponse) error {
			return engine.onMessage(ch, message, response)
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
		go func(ch *CommunicationChannel) {
			err := (*ch).ToHear()
			if err != nil {
				engine.done <- err
			}
		}(&ch)
	}

	return engine, nil
}

func (engine *Engine) Run() error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		// engine.logger.Infoln("Closing all dialogs")
		for _, d := range engine.ActiveDialogs {
			engine.OnContextIsDone(d.Context)
		}
		engine.done <- nil
	}()
	go func() {
		engine.done <- engine.api.Launch()
	}()
	return <-engine.done
}
