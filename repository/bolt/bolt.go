package bolt

import (
	"errors"

	"github.com/boltdb/bolt"
)

var (
	proBucket  string = "Products"
	nameBucket string = "Names"
	errUniq    error  = errors.New("Product name not unique")
	errId      error  = errors.New("Wrong id")
	errName    error  = errors.New("Product dosn't exist")
)

func New() (*bolt.DB, error) {
	db, err := bolt.Open("twiceil.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(proBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(nameBucket))
		if err != nil {
			return err
		}

		return nil
	})

	return db, err
}
