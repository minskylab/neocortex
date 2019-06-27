package neocortex

// Entity define any kind of object or entity
type Entity struct {
	Entity     string                 `json:"entity"`
	Location   []int64                `json:"location"`
	Value      string                 `json:"value"`
	Confidence float64                `json:"confidence"`
	Metadata   map[string]interface{} `json:"metadata"`
}
