package mongo_test

import (
	"booking-books-books-backend/pkg"
	"booking-books-books-backend/pkg/mongo"
	"log"
	"testing"
)

const (
	mongoUrl           = "mongodb://root:password@mongo:27017"
	dbName             = "test_db"
	bookCollectionName = "books"
)

func Test_BookService(t *testing.T) {
	t.Run("CreateBook", createBook_should_insert_book_into_mongo)
}

func createBook_should_insert_book_into_mongo(t *testing.T) {
	//Arrange

	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}

	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()

	bookService := mongo.NewBookService(session.Copy(), dbName, bookCollectionName)

	testTitle := "This is a test book"
	testAuthor := "This is a test author"
	book := root.Book{
		Title:  testTitle,
		Author: testAuthor,
	}

	// Act
	err = bookService.Create(&book)

	// Assert
	if err != nil {
		t.Errorf("Unable to create book: %s", err)
	}
	var results []root.Book
	session.GetCollection(dbName, bookCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect Number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Title != book.Title {
		t.Errorf("Incorrect Title. Expected `%s`, got: `%s`", testTitle, results[0].Title)
	}
}
