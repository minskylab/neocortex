package facebook

import (
	"context"
	"fmt"
	"github.com/bregydoc/neocortex"
	"github.com/k0kubun/pp"
	"github.com/mileusna/facebook-messenger"
	"net/http"
	"strconv"
)

type Channel struct {
	m          *messenger.Messenger
	messageIn  neocortex.MiddleHandler
	newContext neocortex.ContextFabric
	contexts   map[int64]*neocortex.Context
}

type ChannelOptions struct {
	AccessToken string
	VerifyToken string
	PageID      string
}

func (fb *Channel) RegisterMessageEndpoint(handler neocortex.MiddleHandler) error {
	fb.messageIn = handler
	return nil
}

func (fb *Channel) ToHear() error {
	http.Handle("/fb-channel", fb.m)
	fmt.Println("listening on :8080 facebook webhook: /fb-channel")
	return http.ListenAndServe(":8080", nil)
}

func (fb *Channel) GetContextFabric() neocortex.ContextFabric {
	return fb.newContext
}

func (fb *Channel) SetContextFabric(fabric neocortex.ContextFabric) {
	fb.newContext = fabric
}

func NewChannel(options ChannelOptions, fabric ...neocortex.ContextFabric) (*Channel, error) {
	var f neocortex.ContextFabric
	if len(fabric) > 0 {
		f = fabric[0]
	}
	channel := &Channel{
		newContext: f,
		contexts:   map[int64]*neocortex.Context{},
	}

	hook := func(msn *messenger.Messenger, userID int64, m messenger.FacebookMessage) {
		pp.Println(userID)
		uID := strconv.FormatInt(userID, 10)
		pp.Println(uID)
		c, contextExist := channel.contexts[userID]
		if !contextExist {
			c = channel.newContext(context.Background(), uID)
			channel.contexts[userID] = c
		}
		// This is because facebook channel not support entities or intents as input (from messenger chat)
		in := channel.NewInputText(c, m.Text, nil, nil)
		err := channel.messageIn(in, func(out *neocortex.Output) error {
			channel.contexts[userID] = out.Context
			err := decodeOutput(userID, msn, out)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	channel.m = &messenger.Messenger{
		AccessToken:     options.AccessToken,
		VerifyToken:     options.VerifyToken,
		PageID:          options.PageID,
		MessageReceived: hook,
	}

	return channel, nil
}
