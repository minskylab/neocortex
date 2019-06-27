package neocortex

// DialogNode represents a node into complex tree of any dialog
type DialogNode struct {
	Name       string `json:"name"`
	Title      string `json:"title"`
	Conditions string `json:"conditions"`
}
