package neocortex

// Intent define any intent, and intent is like a wish, an intention
type Intent struct {
	Intent     string  `json:"intent"`
	Confidence float64 `json:"confidence"`
}
