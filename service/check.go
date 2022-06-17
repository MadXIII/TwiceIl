package service

import (
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

var bucket string = "Products"

func (c *Check) ToCreate(product *model.Product) (int, int, error) {
	if product.Price < 1 {
		return 0, http.StatusBadRequest, ErrPrice
	}

	id, err := c.repo.Save(bucket, product)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return id, 0, err
}

func (c *Check) ToUpdate(product *model.Product) (int, error) {
	if product.Price < 1 {
		return http.StatusBadRequest, ErrPrice
	}

	err := c.repo.Edit(bucket, product)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func (c *Check) ToDelete(id int) (int, error) {
	if id < 1 {
		return http.StatusBadRequest, ErrId
	}
	if err := c.repo.Delete(bucket, id); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func (c *Check) ToFind(name string) (model.Product, int, error) {
	var product model.Product
	status := 0

	return product, status, nil
}

func (c *Check) ToGet() {
	_, _ = c.repo.Products(1, bucket)
}
