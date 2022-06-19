package bolt

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/madxiii/twiceil/model"
)

var (
	proBucket  string = "Products"
	nameBucket string = "Names"
	errUniq    error  = errors.New("not uniq")
)

type Product struct {
	db *bolt.DB
}

func NewProduct(db *bolt.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Prepare(name string) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(nameBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(proBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(proBucket))
		bytesId := b.Get([]byte(name))
		id := decodeId(bytesId)
		if id > 0 {
			return errUniq
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = p.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameBucket))
		if err != nil {
			return err
		}

		return b.Put([]byte(name), itob(0))
	})

	return err
}

func (p *Product) Save(product *model.Product) (int, error) {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(proBucket))
		if err != nil {
			return err
		}

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		product.Id = int(id)

		buff, err := encodeProd(product)
		if err != nil {
			return err
		}

		return b.Put(itob(product.Id), buff.Bytes())
	})
	if err != nil {
		return 0, err
	}

	err = p.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(nameBucket))
		if err != nil {
			return err
		}

		err = b.Put([]byte(product.Name), itob(product.Id))
		fmt.Println(err)
		return err
	})
	if err != nil {
		return 0, err
	}

	return product.Id, err
}

func (p *Product) Commit(name string, id int) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameBucket))

		return b.Put([]byte(name), itob(id))
	})
	return err
}

// what if name product is duplicate
func (p *Product) Edit(product *model.Product) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(proBucket))
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

// check for wronID
func (p *Product) Delete(id int) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(proBucket))

		err := b.Delete(itob(id))

		return err
	})
	return err
}

func (p *Product) Products() ([]model.Product, error) {
	products := make([]model.Product, 0, 10)
	m := map[string]int{}
	err := p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(nameBucket))
		b.ForEach(func(k, v []byte) error {
			key := decodeKey(k)
			val := decodeId(v)
			fmt.Println(key, val)
			m[key] = val
			// products = append(products, product)
			return nil
		})
		return nil
	})
	fmt.Println(m)
	return products, err
}
