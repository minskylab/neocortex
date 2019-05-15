package neocortex

func (out *Output) ClearResponses() *Output {
	out.Responses = []Response{}
	return out
}

func (out *Output) Clean() *Output {
	out.Responses = []Response{}
	return out
}
