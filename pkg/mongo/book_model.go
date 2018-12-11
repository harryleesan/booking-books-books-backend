package mongo

import (
	"booking-books/books-backend/pkg"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type bookModel struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Title  string
	Author string
}

func bookModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"title", "author"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newBookModel(b *root.Book) *bookModel {
	return &bookModel{
		Title:  b.Title,
		Author: b.Author,
	}
}

func (b *bookModel) toRootBook() *root.Book {
	return &root.Book{
		Id:     b.Id.Hex(),
		Title:  b.Title,
		Author: b.Author,
	}
}
