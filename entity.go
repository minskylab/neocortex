package neocortex

// Entity define any kind of object or entity
type Entity struct {
	Entity     string
	Location   []int64
	Value      string
	Confidence float64
	Metadata   map[string]interface{}
}
