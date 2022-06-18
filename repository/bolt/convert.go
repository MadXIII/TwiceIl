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

func decodeProd(b []byte) model.Product {
	var product model.Product
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	dec.Decode(&product)
	return product
}

func encodeName(name string) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(name)
	return &buff, err
}

func encodeId(id int) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(id)
	return &buff, err
}

func decodeId(b []byte) int {
	var id int
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	dec.Decode(&id)
	return id
}

func itob(id int) []byte {
	return []byte(strconv.FormatInt(int64(id), 10))
}
