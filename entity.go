package neocortex

// Group is a specific kind of entity: it is a group
type Group struct {
	Group    string
	Location []int64
}

// Entity define any kind of object or entity
type Entity struct {
	Entity     string
	Location   []int64
	Value      string
	Confidence float64
	Metadata   map[string]string
	Groups     []Group
}
