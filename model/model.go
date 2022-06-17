package model

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price"`
}

type Delete struct {
	Id int `json:"id"`
}

type Search struct {
	SearchName string `json:"searchName"`
}
