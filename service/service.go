package service

import (
	"github.com/madxiii/twiceil/model"
	"github.com/madxiii/twiceil/repository"
)

type Service struct {
	Checker
}

type Checker interface {
	ToCreate(product *model.Product) (int, int, error)
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Checker: NewCheck(repo),
	}
}
