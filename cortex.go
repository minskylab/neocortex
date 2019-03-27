package neocortex

type OutputResponse func(output *Output) error
type HandleResolver func(in *Input, out *Output, response OutputResponse) error

type CortexMiddleware struct {
	Context             *Context
	Cognitive           CognitiveService
	Channel             CommunicationChannel
	registeredResolvers map[string]*HandleResolver
	genericResolver     *HandleResolver
}

func NewCortex(cognitive CognitiveService, channel CommunicationChannel, c *Context) (*CortexMiddleware, error) {
	middle := &CortexMiddleware{}
	middle.Channel = channel
	middle.Cognitive = cognitive
	middle.registeredResolvers = map[string]*HandleResolver{}
	middle.Context = c
	err := channel.RegisterMessageEndpoint(func(message *Input, response OutputResponse) error {
		// return middle.onMessage(message, response)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return middle, nil
}

//
// func (cortex *CortexMiddleware) onMessage(in *Input, response OutputResponse) error {
// 	out, err := cortex.Cognitive.GetProtoResponse(cortex.Context, in)
// 	if err != nil {
// 		if err == ErrSessionNotExist {
// 			err := response(&Output{
// 				Response: []*ResponseGeneric{{
// 					Text: "I can't to found your session ID",
// 				}},
// 			})
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
//
// 	if len(out.Intents) == 0 {
// 		if cortex.genericResolver != nil {
// 			resolver := *cortex.genericResolver
// 			err = resolver(in, out, response)
// 			if err != nil {
// 				panic(err)
// 			}
// 			return nil
// 		}
// 		err := response(out)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		return nil
// 	}
//
// 	f, ok := cortex.registeredResolvers[out.Intents[0].Intent]
// 	if !ok {
// 		if cortex.genericResolver != nil {
// 			resolver := *cortex.genericResolver
// 			err = resolver(in, out, response)
// 			if err != nil {
// 				panic(err)
// 			}
// 		} else {
// 			err := response(&Output{
// 				Response: []*ResponseGeneric{{
// 					Text: "unimplemented smart response",
// 				}},
// 			})
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
//
// 	} else {
// 		ff := *f
// 		err = ff(in, out, response)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 	}
//
// 	return nil
// }
//
// func (cortex *CortexMiddleware) ResolverAll(handler HandleResolver) {
// 	cortex.genericResolver = &handler
// }
//
// func (cortex *CortexMiddleware) Resolver(entity string, handler HandleResolver) {
// 	if cortex.registeredResolvers == nil {
// 		cortex.registeredResolvers = map[string]*HandleResolver{}
// 	}
// 	cortex.registeredResolvers[entity] = &handler
// }
//
// func (cortex *CortexMiddleware) When(node *DialogNode) {
// 	fmt.Println(node.Title)
// }
