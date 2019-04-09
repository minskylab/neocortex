package boltdb

import "github.com/asdine/storm"

type Repository struct {
	filename string
	db       *storm.DB
}

func New(path string) (*Repository, error) {
	db, err := storm.Open(path)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:       db,
		filename: path,
	}, nil
}
