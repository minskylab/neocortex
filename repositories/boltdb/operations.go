package boltdb

import (
	"github.com/bregydoc/neocortex"
	"github.com/rs/xid"
)

func (repo *Repository) SaveNewDialog(dialog *neocortex.Dialog) (*neocortex.Dialog, error) {
	if dialog.ID == "" {
		dialog.ID = xid.New().String()
	}
	err := repo.db.Save(dialog)
	if err != nil {
		return nil, err
	}
	return repo.GetDialogByID(dialog.ID)
}

func (repo *Repository) GetDialogByID(id string) (*neocortex.Dialog, error) {
	dialog := new(neocortex.Dialog)
	err := repo.db.One("ID", id, dialog)
	if err != nil {
		return nil, err
	}
	return dialog, nil
}

func (repo *Repository) GetAllDialogs() ([]*neocortex.Dialog, error) {
	var allDialogs []*neocortex.Dialog
	err := repo.db.All(&allDialogs)
	if err != nil {
		return nil, err
	}
	return allDialogs, nil
}

func (repo *Repository) GetDialogs(filter neocortex.DialogFilter) ([]*neocortex.Dialog, error) {
	allDialogs, err := repo.GetAllDialogs()
	if err != nil {
		return nil, err
	}
	filteredDialogs := make([]*neocortex.Dialog, 0)
	for _, d := range allDialogs {

		if !filter.From.IsZero() {
			if !d.StartAt.After(filter.From) {
				continue
			}
		}

		if !filter.Until.IsZero() {
			if !d.EndAt.Before(filter.Until) {
				continue
			}
		}

		// if filter.Timezone != "" {
		// 	if d.Context.Person.Timezone != filter.Timezone {
		// 		continue
		// 	}
		// }

		// if filter.SessionID != "" {
		// 	if d.Context.SessionID != filter.SessionID {
		// 		continue
		// 	}
		// }

		// if filter.PersonID != "" {
		// 	if d.Context.Person.ID != filter.PersonID {
		// 		continue
		// 	}
		// }

		filteredDialogs = append(filteredDialogs, d)
	}

	if filter.Limit != 0 {
		return filteredDialogs[:filter.Limit], nil
	}

	return filteredDialogs, nil

}

func (repo *Repository) DeleteDialog(id string) (*neocortex.Dialog, error) {
	d, err := repo.GetDialogByID(id)
	if err != nil {
		return nil, err
	}
	err = repo.db.DeleteStruct(d)
	if err != nil {
		return nil, err
	}

	return d, err
}

func (repo *Repository) UpdateDialog(dialog *neocortex.Dialog) (*neocortex.Dialog, error) {
	err := repo.db.Update(dialog)
	if err != nil {
		return nil, err
	}

	return repo.GetDialogByID(dialog.ID)
}
