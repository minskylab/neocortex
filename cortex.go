package neocortex

import (
	"context"
	"log"
	"time"
)

type OutputResponse func(c *Context, output *Output) error
type HandleResolver func(c *Context, in *Input, out *Output, response OutputResponse) error
type MiddleHandler func(c *Context, message *Input, response OutputResponse) error
type ContextFabric func(ctx context.Context, info PersonInfo) *Context

type Engine struct {
	done                chan error
	cognitive           CognitiveService
	channels            []CommunicationChannel
	registeredResolvers map[CommunicationChannel]map[*Matcher]*HandleResolver
	generalResolver     map[CommunicationChannel]*HandleResolver
	registeredInjection map[CommunicationChannel]map[*Matcher]*InInjection
	generalInjection    map[CommunicationChannel]*InInjection

	Repository    Repository
	ActiveDialogs map[*Context]*Dialog
	api           *API
}

func (engine *Engine) onNewContextCreated(c *Context) {
	log.Println("creating new context: ", c.SessionID)
	engine.ActiveDialogs[c] = newDialog()
}

func (engine *Engine) onContextIsDone(c *Context) {
	for _, ch := range engine.channels {
		ch.CallContextDone(c)
	}
	log.Println("closing context: ", c.SessionID)
	if dialog, ok := engine.ActiveDialogs[c]; ok {
		dialog.EndAt = time.Now()
		if engine.Repository != nil {
			dialog.calcPerformance()
			err := engine.Repository.SaveDialog(dialog)
			if err != nil {
				engine.done <- err
			}
		}
		delete(engine.ActiveDialogs, c)
		log.Println("finally deleting: ", c.SessionID)
	}
}
