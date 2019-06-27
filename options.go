package neocortex

type Option struct {
	Text       string `json:"text"`
	Action     string `json:"action"`
	IsPostBack bool   `json:"is_post_back"`
}

type OptionsResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ItemURL     string    `json:"item_url"`
	Image       string    `json:"image"`
	Options     []*Option `json:"options"`
}
