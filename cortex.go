package neocortex

import (
	"context"
	"time"

	"github.com/asdine/storm"
	"github.com/rs/xid"
)

type OutputResponse func(output *Output) error
type HandleResolver func(in *Input, out *Output, response OutputResponse) error
type MiddleHandler func(message *Input, response OutputResponse) error
type ContextFabric func(ctx context.Context, info PersonInfo) *Context

type Engine struct {
	db                  *storm.DB
	done                chan error
	cognitive           CognitiveService
	channels            []CommunicationChannel
	registeredResolvers map[CommunicationChannel]map[*Matcher]*HandleResolver
	generalResolver     map[CommunicationChannel]*HandleResolver
	sessions            map[string]*Context
	registeredInjection map[CommunicationChannel]map[*Matcher]*InInjection
	generalInjection    map[CommunicationChannel]*InInjection

	Repository    Repository
	ActiveDialogs map[*Context]*Dialog
	api           API
}

func (engine *Engine) OnNewContextCreated(c *Context) {
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
	if _, exist := engine.ActiveDialogs[c]; exist {
		engine.ActiveDialogs[c].EndAt = time.Now()
		if engine.Repository != nil {
			_, err := engine.Repository.SaveNewDialog(engine.ActiveDialogs[c])
			delete(engine.ActiveDialogs, c)
			if err != nil {
				engine.done <- err
			}
		}
	}

}
