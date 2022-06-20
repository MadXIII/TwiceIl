package bolt

import (
	"bytes"

	"github.com/boltdb/bolt"
	"github.com/madxiii/twiceil/model"
)

type Product struct {
	db *bolt.DB
}

func NewProduct(db *bolt.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Save(product *model.Product) (int, error) {
	err := p.db.Update(func(tx *bolt.Tx) error {
		nameB := tx.Bucket([]byte(nameBucket))
		val := nameB.Get([]byte(product.Name))
		if len(val) != 0 {
			return errUniq // fix error
		}
		prodB := tx.Bucket([]byte(proBucket))

		id, err := prodB.NextSequence()
		if err != nil {
			return err
		}

		product.Id = int(id)

		buff, err := encodeProd(product)
		if err != nil {
			return err
		}

		err = prodB.Put(itob(product.Id), buff.Bytes())
		if err != nil {
			return err
		}

		return nameB.Put([]byte(product.Name), itob(product.Id))
	})

	return product.Id, err
}

func (p *Product) Product(id int) (model.Product, error) {
	var product model.Product
	err := p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(proBucket))
		val := b.Get(itob(id))
		var err error
		product, err = decodeProd(val)
		return err
	})
	return product, err
}

func (p *Product) Edit(product *model.Product) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		prodB := tx.Bucket([]byte(proBucket))

		buf := prodB.Get(itob(product.Id))
		prod, err := decodeProd(buf)
		if err != nil {
			return err
		}

		nameB := tx.Bucket([]byte(nameBucket))
		buf = nameB.Get([]byte(prod.Name))
		if !bytes.Equal(buf, itob(product.Id)) {
			return errId
		}
		err = prodB.Delete(itob(product.Id))
		if err != nil {
			return err
		}
		err = nameB.Delete([]byte(prod.Name))
		if err != nil {
			return err
		}

		buff, err := encodeProd(product)
		if err != nil {
			return err
		}
		err = prodB.Put(itob(product.Id), buff.Bytes())
		if err != nil {
			return err
		}
		err = nameB.Put([]byte(product.Name), itob(product.Id))
		return err
	})
	return err
}

func (p *Product) Delete(id int) error {
	err := p.db.Update(func(tx *bolt.Tx) error {
		prodB := tx.Bucket([]byte(proBucket))
		val := prodB.Get(itob(id))

		if len(val) == 0 {
			return errId
		}

		prod, err := decodeProd(val)
		if err != nil {
			return err
		}

		err = prodB.Delete(itob(id))
		if err != nil {
			return err
		}
		nameB := tx.Bucket([]byte(nameBucket))
		return nameB.Delete([]byte(prod.Name))
	})
	return err
}

// Done
func (p *Product) Products() ([]model.Product, error) {
	products := make([]model.Product, 0, 10)
	err := p.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(proBucket))
		err := b.ForEach(func(k, v []byte) error {
			product, err := decodeProd(v)
			if err != nil {
				return err
			}
			products = append(products, product)
			return nil
		})
		return err
	})
	return products, err
}

func (p *Product) Find(name string) (model.Product, error) {
	var product model.Product
	err := p.db.View(func(tx *bolt.Tx) error {
		nameB := tx.Bucket([]byte(nameBucket))
		id := nameB.Get([]byte(name))

		if len(id) == 0 {
			return errName
		}

		prodB := tx.Bucket([]byte(proBucket))
		prod := prodB.Get(id)

		var err error
		product, err = decodeProd(prod)

		return err
	})
	return product, err
}
