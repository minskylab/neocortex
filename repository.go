package neocortex

import "time"

type TimeFramePreset string

const DayPreset TimeFramePreset = "day"
const MonthPreset TimeFramePreset = "month"
const WeekPreset TimeFramePreset = "week"

type TimeFrame struct {
	From     time.Time
	To       time.Time
	PageSize int
	PageNum  int
	Preset   TimeFramePreset
}

type Repository interface {
	SaveDialog(dialog *Dialog) error
	GetDialogByID(id string) (*Dialog, error)
	AllDialogs(frame TimeFrame) ([]*Dialog, error) // to page or not to page?
	DeleteDialog(id string) (*Dialog, error)
	// Dialogs are inmutable, cause they don't have an updater
	DialogsByView(viewID string, frame TimeFrame) ([]*Dialog, error)

	RegisterIntent(intent string) error
	RegisterEntity(entity string) error
	RegisterDialogNode(name string) error
	RegisterContextVar(value string) error
	Intents() []string
	Entities() []string
	DialogNodes() []string
	ContextVars() []string

	SaveView(view *View) error
	GetViewByID(id string) (*View, error)
	FindViewByName(name string) ([]*View, error)
	AllViews() ([]*View, error)
	UpdateView(view *View) error

	SetActionVar(name string, value string) error
	GetActionVar(name string) (string, error)
}
