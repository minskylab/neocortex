package neocortex

import (
	"time"
)

func (engine *Engine) onMessage(channel *CommunicationChannel, in *Input, response OutputResponse) error {

	engine.ActiveDialogs[in.Context].Ins[time.Now()] = in

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

	resolvers, channelIsRegistered := engine.registeredResolvers[*channel]
	if !channelIsRegistered {
		return ErrChannelIsNotRegistered
	}

	exist := false
	for m, resolver := range resolvers {
		if out.Match(&m) {
			if err = (*resolver)(in, out, response); err != nil {
				return err
			}
			engine.ActiveDialogs[in.Context].Outs[time.Now()] = out
			exist = true
		}
	}

	if engine.generalResolver[*channel] != nil && !exist {
		if err = (*engine.generalResolver[*channel])(in, out, response); err != nil {
			return err
		}
		_, activeDialogExist := engine.ActiveDialogs[in.Context]
		if activeDialogExist {
			engine.ActiveDialogs[in.Context].Outs[time.Now()] = out
		}
	}

	return nil
}
