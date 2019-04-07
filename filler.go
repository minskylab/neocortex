package neocortex

import (
	"bytes"
	"text/template"
)

func (out *Output) fill(data interface{}) *Output {
	t := template.New("out")
	buffer := bytes.NewBufferString("")
	newOut := *out
	newOut.Responses = []Response{}
	for _, r := range out.Responses {
		switch r.Type {
		case Text:
			parsed, _ := t.Parse(r.Value.(string))
			_ = parsed.Execute(buffer, data)
			newOut.Responses = append(newOut.Responses, Response{
				Value:    buffer.String(),
				Type:     Text,
				IsTyping: false,
			})
		}
	}
	*out = newOut
	return out
}

func (out *Output) Fill(data interface{}) *Output {
	return out.fill(data)
}
