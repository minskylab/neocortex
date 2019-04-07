package neocortex

import "time"

func (out *Output) AddTextResponse(resp string) *Output {
	if out.Responses == nil {
		out.Responses = []Response{}
	}
	out.Responses = append(out.Responses, Response{
		Type:     Text,
		Value:    resp,
		IsTyping: false,
	})

	return out
}

func (out *Output) AddOptionsResponse(title string, subtitle string, options ...Option) *Output {
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
			Title:       title,
			Description: subtitle,
			Options:     opts,
		},
		IsTyping: false,
	})

	return out
}

func (out *Output) AddListOfOptionsResponse(options []OptionsResponse) *Output {
	if out.Responses == nil {
		out.Responses = []Response{}
	}

	out.Responses = append(out.Responses, Response{
		Type:     Options,
		Value:    options,
		IsTyping: false,
	})

	return out
}

func (out *Output) AddImageResponse(url string) *Output {
	if out.Responses == nil {
		out.Responses = []Response{}
	}

	out.Responses = append(out.Responses, Response{
		Type:     Image,
		Value:    url,
		IsTyping: false,
	})

	return out
}

func (out *Output) AddPauseResponse(duration time.Duration) *Output {
	if out.Responses == nil {
		out.Responses = []Response{}
	}

	out.Responses = append(out.Responses, Response{
		Type:     Pause,
		Value:    duration,
		IsTyping: false,
	})

	return out
}
