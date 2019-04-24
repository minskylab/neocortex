package neocortex

func (out *Output) SetContextVariable(name string, value interface{}) *Output {
	if out.Context.Variables == nil {
		out.Context.Variables = map[string]interface{}{}
	}

	out.Context.Variables[name] = value

	return out
}

func (out *Output) DeleteContextVariable(name string, value interface{}) *Output {
	if out.Context.Variables == nil {
		out.Context.Variables = map[string]interface{}{}
		return out
	}

	delete(out.Context.Variables, name)

	return out
}
