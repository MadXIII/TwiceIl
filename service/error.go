package service

import "errors"

var (
	ErrPrice = errors.New("Price must be positive number")
	ErrId    = errors.New("Wrong Id")
	ErrName  = errors.New("Product name not unique")
)
