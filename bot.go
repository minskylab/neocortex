package neocortex

type Bot struct {
	Name          string `json:"name"`
	FirstGreeting string `json:"first_greeting"`
	Version       string `json:"version"`
	Author        string `json:"author"`
}
