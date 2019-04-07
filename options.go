package neocortex

type Option struct {
	Text       string
	Action     string
	IsPostBack bool
}

type OptionsResponse struct {
	Title    string
	Subtitle string
	ItemURL  string
	Image    string
	Options  []*Option
}
