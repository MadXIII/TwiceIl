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
	Save(bucket string, product *model.Product) (int, error)
	Edit(bucket string, product *model.Product) error
	Delete(bucket string, id int) error
	// Find()
	Products(id int, bucket string) ([]model.Product, error)
}

func New(db *bolt.DB) *Repository {
	return &Repository{
		CRUD: b.NewProduct(db),
	}
}
