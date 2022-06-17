package bolt

import (
	"bytes"
	"encoding/gob"
	"strconv"

	"github.com/madxiii/twiceil/model"
)

func encode(product *model.Product) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(product)
	return &buff, err
}

func decode(b []byte) model.Product {
	var product model.Product
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	dec.Decode(&product)
	return product
}

func itob(id int) []byte {
	return []byte(strconv.FormatInt(int64(id), 10))
}
