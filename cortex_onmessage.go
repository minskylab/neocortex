package neocortex

import (
	"log"
	"time"

	"github.com/k0kubun/pp"
	"github.com/rs/xid"
)

func (engine *Engine) onMessage(channel CommunicationChannel, c *Context, in *Input, response OutputResponse) error {
	pp.Println("context on onMessage: ", c.SessionID)
	inMatched := false
	for matcher, injector := range engine.registeredInjection[channel] {
		if in.Match(nil, matcher) {
			in = (*injector)(c, in)
			inMatched = true
		}
	}

	if engine.generalInjection[channel] != nil && !inMatched {
		f := *engine.generalInjection[channel]
		in = f(c, in)
	}

	if dialog, activeDialogExist := engine.ActiveDialogs[c]; activeDialogExist {
		dialog.LastActivity = time.Now()
		dialog.Contexts[time.Now()] = *c
		dialog.Ins[time.Now()] = *in
	}

	out, err := engine.cognitive.GetProtoResponse(c, in)
	if err != nil {
		if err == ErrSessionNotExist {
			// Creating new context
			log.Println("calling to engine.cognitive.CreateNewContext")
			c1 := engine.cognitive.CreateNewContext(c.Context, c.Person)
			c = c1
			if engine.generalInjection[channel] != nil && !inMatched {
				f := *engine.generalInjection[channel]
				in = f(c, in)
			}

			engine.ActiveDialogs[c] = &Dialog{
				ID:           xid.New().String(),
				LastActivity: time.Now(),
				StartAt:      time.Now(),
				EndAt:        time.Time{},
				Ins:          TimelineInputs{},
				Outs:         TimelineOutputs{},
				Contexts:     TimelineContexts{},
			}

			out, err = engine.cognitive.GetProtoResponse(c, in)
			if err != nil {
				return err
			}

		} else {
			return err
		}
	}

	resolvers, channelIsRegistered := engine.registeredResolvers[channel]
	if !channelIsRegistered {
		return ErrChannelIsNotRegistered
	}

	exist := false
	for m, resolver := range resolvers {
		if out.Match(c, m) {
			if err = (*resolver)(c, in, out, response); err != nil {
				return err
			}

			if dialog, activeDialogExist := engine.ActiveDialogs[c]; activeDialogExist {
				dialog.LastActivity = time.Now()
				dialog.Contexts[time.Now()] = *c
				dialog.Outs[time.Now()] = *out
			}

			exist = true
		}
	}

	if engine.generalResolver[channel] != nil && !exist {
		if err = (*engine.generalResolver[channel])(c, in, out, response); err != nil {
			return err
		}

		if dialog, activeDialogExist := engine.ActiveDialogs[c]; activeDialogExist {
			dialog.LastActivity = time.Now()
			dialog.Contexts[time.Now()] = *c
			dialog.Outs[time.Now()] = *out
		}
	}

	return nil
}
