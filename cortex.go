package neocortex

import (
	"context"
	"github.com/asdine/storm"
	"github.com/sirupsen/logrus"
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
	sessions            map[string]*Context
	Repository          Repository
	ActiveDialogs       map[*Context]*Dialog
}
