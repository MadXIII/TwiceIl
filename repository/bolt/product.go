package bolt

import "github.com/boltdb/bolt"

type Product struct {
	db *bolt.DB
}

func NewProduct(db *bolt.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Create() {
}
