package neocortex

import (
	"context"
	"errors"
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
	Register      map[string]string

	dialogPerformanceFunc func(*Dialog) float64

	secret string
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
			dialog.Performance = engine.dialogPerformanceFunc(dialog)
			err := engine.Repository.SaveDialog(dialog)
			if err != nil {
				engine.done <- err
			}
		}

		delete(engine.ActiveDialogs, c)
		log.Println("finally deleting: ", c.SessionID)
	}
}

// RegisterAdmin you can register new admin for get info purpose
func (engine *Engine) RegisterAdmin(Username, Password string) error {
	if engine.Register != nil {
		engine.Register[Username] = Password
	}
	return nil
}

// getAdmin you can get the password of the admin
func (engine *Engine) getAdmin(Username string) (string, error) {
	if val, ok := engine.Register[Username]; ok {
		return val, nil
	}
	return "", nil
}

// RemoveAdmin let you remove the admin
func (engine *Engine) RemoveAdmin(Username string) error {
	if _, ok := engine.Register[Username]; ok {
		delete(engine.Register, Username)
	}

	return errors.New("no user found")
}
