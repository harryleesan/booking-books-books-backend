package main

import (
	"booking-books-books-backend/pkg/mongo"
	"booking-books-books-backend/pkg/server"
	"log"
	"os"
)

func main() {
	var mongoConn string = "mongodb://root:password@mongo:27017"

	if m := os.Getenv("MONGO_CONN"); m != "" {
		log.Print("MongoDB connection string found. Setting..")
		mongoConn = m
	} else {
		log.Print("MongoDB connection string not set, using default.")
	}

	ms, err := mongo.NewSession(mongoConn)
	if err != nil {
		log.Fatalln("Unable to connect to mongodb")
	}
	defer ms.Close()

	b := mongo.NewBookService(ms.Copy(), "books", "books")
	s := server.NewServer(b)

	s.Start()
}
