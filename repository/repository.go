package repository

import "github.com/boltdb/bolt"

type Repository struct {
	CRUD
}

type CRUD interface {
	Create()
	// Update()
	// Delete()
	// Find()
}

func New(db *bolt.DB) *Repository {
	return &Repository{
		CRUD: bolt.NewProduct(db),
	}
}
