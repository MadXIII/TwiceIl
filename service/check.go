package service

import (
	"net/http"
	"strconv"

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

	id, err := c.repo.Save(product)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return id, 0, err
}

func (c *Check) ToProduct(param string) (model.Product, int, error) {
	var product model.Product

	id, err := strconv.Atoi(param)
	if err != nil {
		return product, http.StatusBadRequest, ErrId
	}

	product, err = c.repo.Product(id)
	if err != nil {
		return product, http.StatusInternalServerError, err
	}

	return product, 0, nil
}

func (c *Check) ToUpdate(product *model.Product) (int, error) {
	if product.Price < 1 {
		return http.StatusBadRequest, ErrPrice
	}

	err := c.repo.Edit(product)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func (c *Check) ToDelete(id int) (int, error) {
	if id < 1 {
		return http.StatusBadRequest, ErrId
	}
	if err := c.repo.Delete(id); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func (c *Check) ToFind(name string) (model.Product, error) {
	var product model.Product
	if len(name) == 0 {
		return product, ErrEmpty
	}

	product, err := c.repo.Find(name)

	return product, err
}

func (c *Check) ToProducts() ([]model.Product, int, error) {
	products, err := c.repo.Products()

	return products, http.StatusInternalServerError, err
}
