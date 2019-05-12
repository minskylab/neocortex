package neocortex

func (out *Output) ClearResponses() *Output {
	out.Responses = []Response{}
	return out
}

func (out *Output) Clean() *Output {
	out.Responses = []Response{}
	return out
}

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

func (out *Output) GetStringContextVariable(name string, _default string) string {
	if out.Context.Variables == nil {
		out.Context.Variables = map[string]interface{}{}
		return _default
	}

	value, isOk := out.Context.Variables[name].(string)
	if !isOk {
		return _default
	}

	return value
}

func (out *Output) GetIntContextVariable(name string, _default int) int {
	if out.Context.Variables == nil {
		out.Context.Variables = map[string]interface{}{}
		return _default
	}

	value, isOk := out.Context.Variables[name].(int)
	if !isOk {
		return _default
	}

	return value
}

func (out *Output) GetFloat64ContextVariable(name string, _default float64) float64 {
	if out.Context.Variables == nil {
		out.Context.Variables = map[string]interface{}{}
		return _default
	}

	value, isOk := out.Context.Variables[name].(float64)
	if !isOk {
		return _default
	}

	return value
}
