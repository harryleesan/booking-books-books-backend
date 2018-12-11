package main

import (
	"booking-books/books-backend/pkg/mongo"
	"booking-books/books-backend/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("mongodb://root:password@localhost:27017")
	if err != nil {
		log.Fatalln("Unable to connect to mongodb")
	}
	defer ms.Close()

	b := mongo.NewBookService(ms.Copy(), "books", "books")
	s := server.NewServer(b)

	s.Start()
}
