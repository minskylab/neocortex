package neocortex

type ViewStyle string

const Line ViewStyle = "line"
const Bars ViewStyle = "bars"
const Pie ViewStyle = "pie"
const Map ViewStyle = "map"

type ViewClassType string

const EntityClass ViewClassType = "entity"
const IntentClass ViewClassType = "intent"
const DialogNodeClass ViewClassType = "node"
const ContextVarClass ViewClassType = "context"

type ViewClass struct {
	Type  ViewClassType `json:"type"`
	Value string        `json:"value"`
}

type ActionVarType string

const TextActionVar ActionVarType = "text"

type View struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Styles   []ViewStyle `json:"styles"`
	Classes  []ViewClass `json:"classes"`
	Children []*View     `json:"children"`
}

type ActionVariable struct {
	Name  string        `json:"name"`
	Type  ActionVarType `json:"type"`
	Value []byte        `json:"value"`
}

type UsersSummary struct {
	News       int64 `json:"news"`
	Recurrents int64 `json:"recurrents"`
}

type Summary struct {
	TotalDialogs    int64                   `json:"total_dialogs"`
	NewUsers        int64                   `json:"new_users"`
	RecurrentUsers  int64                   `json:"recurrent_users"`
	UsersByTimezone map[string]UsersSummary `json:"users_by_timezone"`
}
