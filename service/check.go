package service

import "github.com/madxiii/twiceil/repository"

type Check struct {
	repo *repository.Repository
}

func NewCheck(repo *repository.Repository) *Check {
	return &Check{repo: repo}
}

func (c *Check) Create() {
}
