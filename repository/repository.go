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
	Save(product *model.Product) (int, error)
	Edit(product *model.Product) error
	Delete(id int) error
	Product(id int) (model.Product, error)
	Find(name string) (model.Product, error)
	Products() ([]model.Product, error)
}

func New(db *bolt.DB) *Repository {
	return &Repository{
		CRUD: b.NewProduct(db),
	}
}
