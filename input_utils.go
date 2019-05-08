package neocortex

func (in *Input) SetContextVariable(name string, value interface{}) *Input {
	if in.Context.Variables == nil {
		in.Context.Variables = map[string]interface{}{}
	}

	in.Context.Variables[name] = value

	return in
}

func (in *Input) DeleteContextVariable(name string, value interface{}) *Input {
	if in.Context.Variables == nil {
		in.Context.Variables = map[string]interface{}{}
		return in
	}

	delete(in.Context.Variables, name)

	return in
}
