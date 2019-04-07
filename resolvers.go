package neocortex

func (engine *Engine) ResolveAny(channel CommunicationChannel, handler HandleResolver) {
	if engine.generalResolver == nil {
		engine.generalResolver = map[CommunicationChannel]*HandleResolver{}
	}
	engine.generalResolver[channel] = &handler
}

func (engine *Engine) Resolve(channel CommunicationChannel, matcher Matcher, handler HandleResolver) {
	if engine.registeredResolvers == nil {
		engine.registeredResolvers = map[CommunicationChannel]map[Matcher]*HandleResolver{}
	}
	if engine.registeredResolvers[channel] == nil {
		engine.registeredResolvers[channel] = map[Matcher]*HandleResolver{}
	}
	engine.registeredResolvers[channel][matcher] = &handler
}

func (engine *Engine) ResolveMany(channels []CommunicationChannel, matcher Matcher, handler HandleResolver) {
	for _, ch := range channels {
		engine.Resolve(ch, matcher, handler)
	}
}

func (engine *Engine) ResolveManyAny(channels []CommunicationChannel, handler HandleResolver) {
	for _, ch := range channels {
		engine.ResolveAny(ch, handler)
	}
}
