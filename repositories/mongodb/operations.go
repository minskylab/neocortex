package mongodb

import (
	"context"

	"github.com/bregydoc/neocortex"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *Repository) SaveNewDialog(dialog *neocortex.Dialog) (*neocortex.Dialog, error) {
	if dialog.ID == "" {
		dialog.ID = xid.New().String()
	}
	_, err := repo.dialogsCollection.InsertOne(context.Background(), dialogToDocument(dialog))
	if err != nil {
		return nil, err
	}
	return repo.GetDialogByID(dialog.ID)
}

func (repo *Repository) GetDialogByID(id string) (*neocortex.Dialog, error) {
	result := repo.dialogsCollection.FindOne(context.Background(), bson.M{"id": id})
	if result.Err() != nil {
		return nil, result.Err()
	}
	dialog := new(DialogDocument)
	err := result.Decode(dialog)
	if err != nil {
		return nil, err
	}
	return documentToDialog(dialog), nil
}

func (repo *Repository) GetAllDialogs() ([]*neocortex.Dialog, error) {
	cursor, err := repo.dialogsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	dialogs := make([]*neocortex.Dialog, 0)
	for cursor.Next(context.Background()) {
		if cursor.Err() != nil {
			return nil, cursor.Err()
		}
		dialog := new(DialogDocument)
		err := cursor.Decode(dialog)
		if err != nil {
			return nil, err
		}
		dialogs = append(dialogs, documentToDialog(dialog))
	}

	return dialogs, nil
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
	dialog, err := repo.GetDialogByID(id)
	if err != nil {
		return nil, err
	}
	_, err = repo.dialogsCollection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return nil, err
	}
	return dialog, nil
}

func (repo *Repository) UpdateDialog(dialog *neocortex.Dialog) (*neocortex.Dialog, error) {

	_, err := repo.dialogsCollection.UpdateOne(context.Background(), bson.M{"id": dialog.ID}, dialogToDocument(dialog))
	if err != nil {
		return nil, err
	}

	return repo.GetDialogByID(dialog.ID)
}
