package repository

import (
	"github.com/boltdb/bolt"
	"github.com/madxiii/twiceil/model"
	b "github.com/madxiii/twiceil/repository/bolt"
)

type Repository struct {
	CRUD
}

type CRUD interface {
	Prepare(name string) error
	Save(product *model.Product) (int, error)
	Commit(name string, id int) error
	Edit(product *model.Product) error
	Delete(id int) error
	// Find()
	Products() ([]model.Product, error)
}

func New(db *bolt.DB) *Repository {
	return &Repository{
		CRUD: b.NewProduct(db),
	}
}
