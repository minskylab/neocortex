package neocortex

func (out *Output) AddTextResponse(resp string) {
	if out.Responses == nil {
		out.Responses = []Response{}
	}
	out.Responses = append(out.Responses, Response{
		Type:     Text,
		Value:    resp,
		IsTyping: false,
	})
}

func (out *Output) AddOptionsResponse(title string, subtitle string, options ...Option) {
	if out.Responses == nil {
		out.Responses = []Response{}
	}
	opts := make([]*Option, 0)
	for _, o := range options {
		opts = append(opts, &o)
	}
	out.Responses = append(out.Responses, Response{
		Type: Options,
		Value: OptionsResponse{
			Title:    title,
			Subtitle: subtitle,
			Options:  opts,
		},
		IsTyping: false,
	})
}

func (out *Output) AddImageResponse(url string) {
	if out.Responses == nil {
		out.Responses = []Response{}
	}

	out.Responses = append(out.Responses, Response{
		Type:     Image,
		Value:    url,
		IsTyping: false,
	})
}
