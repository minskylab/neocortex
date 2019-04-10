package neocortex

import (
	"reflect"
)

func (engine *Engine) onMessage(channel *CommunicationChannel, in *Input, response OutputResponse) error {
	entry := engine.logger.WithField("from", reflect.ValueOf(channel).Type())
	entry.Debug("new message in")
	out, err := engine.cognitive.GetProtoResponse(in)
	if err != nil {
		entry.WithField("error", err).Debug("error launched")
		if err == ErrSessionNotExist {
			// TODO: Check above later, it's so strange
			f := (*channel).GetContextFabric()
			f(*in.Context.Context, in.Context.Person)
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
			exist = true
		}
	}

	if engine.generalResolver[*channel] != nil && !exist {
		if err = (*engine.generalResolver[*channel])(in, out, response); err != nil {
			return err
		}
	}

	return nil
}
