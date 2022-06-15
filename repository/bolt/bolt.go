package bolt

import "github.com/boltdb/bolt"

func New() (*bolt.DB, error) {
	db, err := bolt.Open("twiceil.db", 0600, nil)

	return db, err
}
