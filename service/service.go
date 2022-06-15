package service

import "github.com/madxiii/twiceil/repository"

type Service struct {
	Checker
}

type Checker interface {
	Create()
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Checker: NewCheck(repo),
	}
}
