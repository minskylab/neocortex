package facebook

import (
	"context"
	"fmt"
	"log"
	"strconv"

	neo "github.com/bregydoc/neocortex"
	"github.com/bregydoc/neocortex/channels/facebook/messenger"
)

func NewChannel(options ChannelOptions, fabric ...neo.ContextFabric) (*Channel, error) {
	fb := &Channel{
		contexts: map[int64]*neo.Context{},
	}

	if len(fabric) > 0 {
		f := fabric[0]
		fb.newContext = f
	}

	hook := func(msn *messenger.Messenger, user messenger.UserInfo, m messenger.FacebookMessage) {
		uID := strconv.FormatInt(user.ID, 10)
		tz := fmt.Sprintf("%d", int(user.Timezone))
		c, contextExist := fb.contexts[user.ID]
		if !contextExist {

			c = fb.newContext(context.Background(), neo.PersonInfo{
				ID:       uID,
				Timezone: tz,
				Name:     user.Name,
				Locale:   user.Locale,
				Picture:  user.ProfilePic,
			})

			for _, call := range fb.newContextCallbacks {
				(*call)(c)
			}
			fb.contexts[user.ID] = c
		}
		// This is because facebook channel not support entities or intents as input (from messenger chat)
		in := fb.NewInputText(c, m.Text, nil, nil)
		err := fb.messageIn(in, func(out *neo.Output) error {
			fb.contexts[user.ID] = out.Context
			err := decodeOutput(user.ID, msn, out)

			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
	}

	postbackHook := func(msn *messenger.Messenger, user messenger.UserInfo, p messenger.FacebookPostback) {
		text := p.Payload

		uID := strconv.FormatInt(user.ID, 10)
		tz := fmt.Sprintf("%d", int(user.Timezone))
		c, contextExist := fb.contexts[user.ID]
		if !contextExist {
			c = fb.newContext(context.Background(), neo.PersonInfo{
				ID:       uID,
				Timezone: tz,
				Name:     user.Name,
			})

			for _, call := range fb.newContextCallbacks {
				(*call)(c)
			}
			fb.contexts[user.ID] = c
		}
		// This is because facebook channel not support entities or intents as input (from messenger chat)
		in := fb.NewInputText(c, text, nil, nil)
		err := fb.messageIn(in, func(out *neo.Output) error {
			fb.contexts[user.ID] = out.Context
			err := decodeOutput(user.ID, msn, out)

			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
	}

	fb.m = &messenger.Messenger{
		AccessToken:      options.AccessToken,
		VerifyToken:      options.VerifyToken,
		PageID:           options.PageID,
		MessageReceived:  hook,
		PostbackReceived: postbackHook,
	}

	return fb, nil
}
