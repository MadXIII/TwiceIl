package bolt

import (
	"bytes"

	"github.com/boltdb/bolt"
)

type Product struct {
	db *bolt.DB
}

func NewProduct(db *bolt.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Created(buffer *bytes.Buffer) (int, error) {
	return 0, nil
}
