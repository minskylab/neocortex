package neocortex

import "time"

type DialogFilter struct {
	PersonID  string
	From      time.Time
	Until     time.Time
	SessionID string
	Timezone  string
	Limit     int64
}

type Repository interface {
	SaveNewDialog(dialog *Dialog) (*Dialog, error)
	GetDialogByID(id string) (*Dialog, error)
	GetAllDialogs() ([]*Dialog, error)
	GetDialogs(filter DialogFilter) ([]*Dialog, error)
	DeleteDialog(id string) (*Dialog, error)
	UpdateDialog(dialog *Dialog) (*Dialog, error)
}
