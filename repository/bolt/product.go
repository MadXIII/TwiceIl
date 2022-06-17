package bolt

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/madxiii/twiceil/model"
)

type Product struct {
	db *bolt.DB
}

func NewProduct(db *bolt.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Save(bucket string, product *model.Product) (int, error) {
	var prodId int
	err := p.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		prodId = int(id)
		product.Id = int(id)

		buff, err := encode(product)
		if err != nil {
			return err
		}

		return b.Put(itob(product.Id), buff.Bytes())
	})
	return prodId, err
}

func (p *Product) Edit(bucket string, product *model.Product) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		buff, err := encode(product)
		if err != nil {
			return err
		}

		return b.Put(itob(product.Id), buff.Bytes())
	})
	return err
}

func (p *Product) Delete(bucket string, id int) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		err := b.Delete(itob(id))

		return err
	})
	return err
}

func (p *Product) Products(key int, bucket string) ([]model.Product, error) {
	var products []model.Product
	p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		bytes := b.Get(itob(key))
		fmt.Println(decode(bytes))
		return nil
	})

	return products, nil
}
