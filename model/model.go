package model

type Product struct {
	Id    int    `json:"-"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price"`
}

type Search struct {
	SearchName string `json:"searchName"`
}
