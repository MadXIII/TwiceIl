package main

import (
	"log"

	"github.com/madxiii/twiceil/handler"
	"github.com/madxiii/twiceil/repository"
	"github.com/madxiii/twiceil/repository/bolt"
	"github.com/madxiii/twiceil/service"
)

func main() {
	db, err := bolt.New()
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}
	defer db.Close()

	repo := repository.New(db)
	services := service.New(repo)
	handlers := handler.New(services).Routes()

	serv := new(handler.Server)

	if err := serv.Run(":8282", handlers); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
