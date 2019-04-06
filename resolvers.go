package neocortex

func (cortex *Engine) ResolveAny(channel CommunicationChannel, handler HandleResolver) {
	if cortex.generalResolver == nil {
		cortex.generalResolver = map[CommunicationChannel]*HandleResolver{}
	}
	cortex.generalResolver[channel] = &handler
}

func (cortex *Engine) Resolve(channel CommunicationChannel, matcher Matcher, handler HandleResolver) {
	if cortex.registeredResolvers == nil {
		cortex.registeredResolvers = map[CommunicationChannel]map[Matcher]*HandleResolver{}
	}
	if cortex.registeredResolvers[channel] == nil {
		cortex.registeredResolvers[channel] = map[Matcher]*HandleResolver{}
	}
	cortex.registeredResolvers[channel][matcher] = &handler
}

func (cortex *Engine) ResolveMany(channels []CommunicationChannel, matcher Matcher, handler HandleResolver) {
	for _, ch := range channels {
		cortex.Resolve(ch, matcher, handler)
	}
}

func (cortex *Engine) ResolveManyAny(channels []CommunicationChannel, handler HandleResolver) {
	for _, ch := range channels {
		cortex.ResolveAny(ch, handler)
	}
}
