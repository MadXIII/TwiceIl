package bolt

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/madxiii/twiceil/model"
)

var errUniq = errors.New("not uniq")

type Product struct {
	db *bolt.DB
}

func NewProduct(db *bolt.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Prepare(bucket string, name string) error {
	err := p.db.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			fmt.Println(err)
			return err
		}
		bytesId := b.Get([]byte(name))
		fmt.Println("IDBYTES", bytesId)
		id := decodeId(bytesId)
		if id < 1 {
			return errUniq
		}

		return b.Put([]byte(name), bytesId)
	})
	fmt.Println("Prepare", err)
	return err
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

		prodId = int(id) // ubrat'?
		product.Id = int(id)

		buff, err := encodeProd(product)
		if err != nil {
			return err
		}

		return b.Put(itob(product.Id), buff.Bytes())
	})
	return prodId, err
}

func (p *Product) Commit(bucket string, name string, id int) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		buffName, err := encodeName(name)
		if err != nil {
			return err
		}

		buffId, err := encodeId(id)
		if err != nil {
			return err
		}

		return b.Put(buffName.Bytes(), buffId.Bytes())
	})
	return err
}

func (p *Product) Edit(bucket string, product *model.Product) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		buff, err := encodeProd(product)
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
	products := make([]model.Product, 0, 10)
	err := p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		lastKey := int(b.Sequence())
		for i := 1; i <= lastKey; i++ {
			bytes := b.Get(itob(i))
			products = append(products, decodeProd(bytes))
		}
		return nil
	})

	return products, err
}
