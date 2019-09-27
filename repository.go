package neocortex

import "time"

type TimeFramePreset string

const DayPreset TimeFramePreset = "day"
const MonthPreset TimeFramePreset = "month"
const WeekPreset TimeFramePreset = "week"
const YearPreset TimeFramePreset = "year"

type TimeFrame struct {
	From     time.Time
	To       time.Time
	Preset   TimeFramePreset
	PageSize int
	PageNum  int
}

type Repository interface {
	SaveDialog(dialog *Dialog) error
	GetDialogByID(id string) (*Dialog, error)
	AllDialogs(frame TimeFrame) ([]*Dialog, error) // to page or not to page?
	DeleteDialog(id string) (*Dialog, error)

	// Dialogs are inmutable, cause it doesn't have an updater
	DialogsByView(viewID string, frame TimeFrame) ([]*Dialog, error)
	Summary(frame TimeFrame) (*Summary, error)

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
	DeleteView(id string) (*View, error)

	SetActionVar(name string, value string) error
	GetActionVar(name string) (string, error)
}
