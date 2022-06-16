package repository

import (
	"bytes"

	"github.com/boltdb/bolt"
	b "github.com/madxiii/twiceil/repository/bolt"
)

type Repository struct {
	CRUD
}

type CRUD interface {
	Created(bytes *bytes.Buffer) (int, error)
	// Update()
	// Delete()
	// Find()
}

func New(db *bolt.DB) *Repository {
	return &Repository{
		CRUD: b.NewProduct(db),
	}
}
