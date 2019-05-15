package neocortex

type InInjection func(c *Context, in *Input) *Input

func (engine *Engine) InjectAll(channel CommunicationChannel, middle InInjection) {
	if engine.generalInjection == nil {
		engine.generalInjection = map[CommunicationChannel]*InInjection{}
	}
	engine.generalInjection[channel] = &middle
}

func (engine *Engine) Inject(channel CommunicationChannel, matcher *Matcher, middle InInjection) {
	if engine.registeredInjection == nil {
		engine.registeredInjection = map[CommunicationChannel]map[*Matcher]*InInjection{}
	}

	if engine.registeredInjection[channel] == nil {
		engine.registeredInjection[channel] = map[*Matcher]*InInjection{}
	}

	engine.registeredInjection[channel][matcher] = &middle
}
