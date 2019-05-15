package neocortex

import (
	"context"
	"log"
	"time"

	"github.com/asdine/storm"
	"github.com/rs/xid"
)

type OutputResponse func(c *Context, output *Output) error
type HandleResolver func(c *Context, in *Input, out *Output, response OutputResponse) error
type MiddleHandler func(c *Context, message *Input, response OutputResponse) error
type ContextFabric func(ctx context.Context, info PersonInfo) *Context

type Engine struct {
	db                  *storm.DB
	done                chan error
	cognitive           CognitiveService
	channels            []CommunicationChannel
	registeredResolvers map[CommunicationChannel]map[*Matcher]*HandleResolver
	generalResolver     map[CommunicationChannel]*HandleResolver
	registeredInjection map[CommunicationChannel]map[*Matcher]*InInjection
	generalInjection    map[CommunicationChannel]*InInjection

	Repository    Repository
	ActiveDialogs map[*Context]*Dialog
	api           API
}

func (engine *Engine) OnNewContextCreated(c *Context) {
	log.Println("creating new context: ", c.SessionID)
	engine.ActiveDialogs[c] = &Dialog{
		ID:      xid.New().String(),
		Context: c,
		StartAt: time.Now(),
		EndAt:   time.Time{},
		Ins:     TimelineInputs{},
		Outs:    TimelineOutputs{},
	}
}

func (engine *Engine) OnContextIsDone(c *Context) {
	for _, ch := range engine.channels {
		ch.CallContextDone(c)
	}
	log.Printf("%v\n", engine.ActiveDialogs)
	log.Println("closing context: ", c.SessionID)
	log.Println("context to delete:", c)
	engine.ActiveDialogs[c].EndAt = time.Now()
	if engine.Repository != nil {
		_, err := engine.Repository.SaveNewDialog(engine.ActiveDialogs[c])
		if err != nil {
			engine.done <- err
		}
	}
	delete(engine.ActiveDialogs, c)
	log.Println("finally deleting: ", c.SessionID)
}
