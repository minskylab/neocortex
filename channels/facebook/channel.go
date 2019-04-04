package facebook

import (
	"fmt"
	"github.com/bregydoc/neocortex"
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

func CreateNewFacebookChannel(accessToken, verifyToken, pageID string) (*Channel, error) {
	channel := &Channel{}
	hook := func(msn *messenger.Messenger, userID int64, m messenger.FacebookMessage) {
		c, contextAlreadyExist := channel.contexts[userID]
		if !contextAlreadyExist {
			c = channel.newContext(strconv.Itoa(int(userID)))
		}
		// This is because facebook channel not support entities or intents as input (from messenger chat)
		in := channel.NewInputText(c, m.Text, nil, nil)
		err := channel.messageIn(in, func(output neocortex.Output) error {
			channel.contexts[userID] = output.Context()
			in = channel.NewInputText(channel.contexts[userID], m.Text, output.Entities(), output.Intents())
			for _, r := range output.Responses() {
				_, err := msn.SendTextMessage(userID, r.Value().(string))
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	channel.m = &messenger.Messenger{
		AccessToken:     accessToken,
		VerifyToken:     verifyToken,
		PageID:          pageID,
		MessageReceived: hook,
	}

	return channel, nil
}
