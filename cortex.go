package neocortex

import (
	"context"
	"errors"
	"github.com/asdine/storm"
	"github.com/sirupsen/logrus"
	"reflect"
)

type OutputResponse func(output *Output) error
type HandleResolver func(in *Input, out *Output, response OutputResponse) error
type MiddleHandler func(message *Input, response OutputResponse) error
type ContextFabric func(ctx context.Context, info PersonInfo) *Context

type Engine struct {
	logger              *logrus.Logger
	db                  *storm.DB
	done                chan error
	cognitive           CognitiveService
	channels            []CommunicationChannel
	registeredResolvers map[CommunicationChannel]map[Matcher]*HandleResolver
	generalResolver     map[CommunicationChannel]*HandleResolver
}

func (engine *Engine) onMessage(channel CommunicationChannel, in *Input, response OutputResponse) error {
	entry := engine.logger.WithField("from", reflect.ValueOf(channel).Type())
	entry.Debug("new message in")
	out, err := engine.cognitive.GetProtoResponse(in)
	if err != nil {
		entry.WithField("error", err).Debug("error launched")
		if err == ErrSessionNotExist {
			// TODO: Check above later, it's so strange
			f := channel.GetContextFabric()
			f(*in.Context.Context, in.Context.Person)
		} else {
			return err
		}
	}

	resolvers, channelIsRegistered := engine.registeredResolvers[channel]
	if !channelIsRegistered {
		return errors.New("channel not exist on this engine instance")
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

	if engine.generalResolver != nil && !exist {
		if err = (*engine.generalResolver[channel])(in, out, response); err != nil {
			return err
		}
	}

	return nil
}
