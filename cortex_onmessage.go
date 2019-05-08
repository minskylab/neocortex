package neocortex

import (
	"fmt"
	"github.com/k0kubun/pp"
	"time"
)

func (engine *Engine) onMessage(channel *CommunicationChannel, in *Input, response OutputResponse) error {
	pp.Println("IN: ", in)
	_, activeDialogExist := engine.ActiveDialogs[in.Context]
	if activeDialogExist {
		engine.ActiveDialogs[in.Context].Ins[time.Now()] = in
	}
	out, err := engine.cognitive.GetProtoResponse(in)
	if err != nil {
		if err == ErrSessionNotExist {
			f := (*channel).GetContextFabric()
			// Creating new context
			c := f(*in.Context.Context, in.Context.Person)
			in.Context = c
			out, err = engine.cognitive.GetProtoResponse(in)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	pp.Println("OUT: ", out)

	resolvers, channelIsRegistered := engine.registeredResolvers[*channel]
	if !channelIsRegistered {
		return ErrChannelIsNotRegistered
	}
	
	fmt.Printf("channel: %p\n", channel)
	pp.Println(len(resolvers))
	exist := false
	for m, resolver := range resolvers {
		if out.Match(m) {
			if err = (*resolver)(in, out, response); err != nil {
				return err
			}

			if _, ok := engine.ActiveDialogs[in.Context]; ok {
				engine.ActiveDialogs[in.Context].Outs[time.Now()] = out
			}

			exist = true
		}
	}

	if engine.generalResolver[*channel] != nil && !exist {
		if err = (*engine.generalResolver[*channel])(in, out, response); err != nil {
			return err
		}
		if _, ok := engine.ActiveDialogs[in.Context]; ok {
			engine.ActiveDialogs[in.Context].Outs[time.Now()] = out
		}
	}

	return nil
}
