package bolt

import (
	"bytes"
	"encoding/gob"
	"strconv"

	"github.com/madxiii/twiceil/model"
)

func encodeProd(product *model.Product) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(product)
	return &buff, err
}

func decodeProd(b []byte) (model.Product, error) {
	var product model.Product
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&product)
	return product, err
}

func itob(id int) []byte {
	return []byte(strconv.FormatInt(int64(id), 10))
}
