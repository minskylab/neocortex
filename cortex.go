package neocortex

import (
	"context"
	"errors"
	"github.com/asdine/storm"
)

type OutputResponse func(output *Output) error
type HandleResolver func(in *Input, out *Output, response OutputResponse) error
type MiddleHandler func(message *Input, response OutputResponse) error
type ContextFabric func(ctx context.Context, userID string) *Context

type Engine struct {
	db                  *storm.DB
	done                chan error
	cognitive           CognitiveService
	channels            []CommunicationChannel
	registeredResolvers map[CommunicationChannel]map[Matcher]*HandleResolver
	generalResolver     map[CommunicationChannel]*HandleResolver
}

func (cortex *Engine) onMessage(channel CommunicationChannel, in *Input, response OutputResponse) error {
	out, err := cortex.cognitive.GetProtoResponse(in)
	if err != nil {
		if err == ErrSessionNotExist {
			// TODO: Check above later, it's so strange
			f := channel.GetContextFabric()
			f(*in.Context.Context, in.Context.UserID)
		} else {
			return err
		}
	}

	resolvers, channelIsRegistered := cortex.registeredResolvers[channel]
	if !channelIsRegistered {
		return errors.New("channel not exist on this neocortex instance")
	}

	exist := false
	for m, resolver := range resolvers {
		if match(out, &m) {
			if err = (*resolver)(in, out, response); err != nil {
				return err
			}
			exist = true
		}
	}

	if cortex.generalResolver != nil && !exist {
		if err = (*cortex.generalResolver[channel])(in, out, response); err != nil {
			return err
		}
	}

	return nil
}
