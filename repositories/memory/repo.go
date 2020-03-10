package memory

import (
	"github.com/minskylab/neocortex"
)

type InMemoryRepo struct {
	dialogs []*neocortex.Dialog
}

func (m *InMemoryRepo) SaveDialog(dialog *neocortex.Dialog) error {
	m.dialogs = append(m.dialogs, dialog)
	return nil
}
