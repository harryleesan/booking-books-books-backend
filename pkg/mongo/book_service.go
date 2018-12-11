package mongo

import (
	"booking-books/books-backend/pkg"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookService struct {
	collection *mgo.Collection
}

func NewBookService(session *Session, dbName string, collectionName string) *BookService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(bookModelIndex())
	return &BookService{
		collection,
	}
}

func (p *BookService) CreateBook(b *root.Book) error {
	book := newBookModel(b)
	return p.collection.Insert(&book)
}

func (p *BookService) GetByTitle(title string) (*root.Book, error) {
	model := bookModel{}
	err := p.collection.Find(bson.M{"title": title}).One(&model)
	return model.toRootBook(), err
}

func (p *BookService) GetByAuthor(author string) (*root.Book, error) {
	model := bookModel{}
	err := p.collection.Find(bson.M{"author": author}).One(&model)
	return model.toRootBook(), err
}
