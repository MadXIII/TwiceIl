package service

import (
	"bytes"
	"encoding/gob"
	"net/http"

	"github.com/madxiii/twiceil/model"
	"github.com/madxiii/twiceil/repository"
)

type Check struct {
	repo *repository.Repository
}

func NewCheck(repo *repository.Repository) *Check {
	return &Check{repo: repo}
}

func (c *Check) ToCreate(product *model.Product) (int, int, error) {
	if product.Price < 1 {
		return 0, http.StatusBadRequest, ErrPrice
	}

	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	enc.Encode(product)

	id, err := c.repo.Created(&buff)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return 0, id, err
}
