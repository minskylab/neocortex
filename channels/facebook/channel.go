package facebook

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/facebook/messenger"
)

type Channel struct {
	m                    *messenger.Messenger
	messageIn            neo.MiddleHandler
	newContext           neo.ContextFabric
	contexts             map[int64]*neo.Context
	newContextCallbacks  []*func(c *neo.Context)
	doneContextCallbacks []*func(c *neo.Context)
}

type ChannelOptions struct {
	AccessToken string
	VerifyToken string
	PageID      string
}

func (fb *Channel) RegisterMessageEndpoint(handler neo.MiddleHandler) error {
	fb.messageIn = handler
	return nil
}

func (fb *Channel) ToHear() error {
	http.Handle("/fb-channel", fb.m)
	fmt.Println("listening on :8080 facebook webhook: /fb-channel")
	return http.ListenAndServe(":8080", nil)
}

func (fb *Channel) GetContextFabric() neo.ContextFabric {
	return fb.newContext
}

func (fb *Channel) SetContextFabric(fabric neo.ContextFabric) {
	fb.newContext = fabric
}

func (fb *Channel) OnNewContextCreated(callback func(c *neo.Context)) {
	if fb.newContextCallbacks == nil {
		fb.newContextCallbacks = []*func(c *neo.Context){}
	}
	fb.newContextCallbacks = append(fb.newContextCallbacks, &callback)
}

func (fb *Channel) OnContextIsDone(callback func(c *neo.Context)) {
	if fb.doneContextCallbacks == nil {
		fb.doneContextCallbacks = []*func(c *neo.Context){}
	}
	fb.doneContextCallbacks = append(fb.doneContextCallbacks, &callback)
}

func (fb *Channel) CallContextDone(c *neo.Context) {
	log.Println("fb, CallContextDone: ", c.SessionID)
	id, err := strconv.ParseInt(c.Person.ID, 10, 64)
	if err == nil {
		delete(fb.contexts, id)
	}
}
