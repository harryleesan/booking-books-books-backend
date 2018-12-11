package mongo

import (
	"booking-books/books-backend/pkg"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
	err := p.collection.Find(bson.M{"title": bson.RegEx{fmt.Sprintf(".*%s.*", title), "i"}}).One(&model)
	return model.toRootBook(), err
}

func (p *BookService) GetByAuthor(author string) (*root.Book, error) {
	model := bookModel{}
	err := p.collection.Find(bson.M{"author": bson.RegEx{fmt.Sprintf(".*%s.*", author), "i"}}).One(&model)
	return model.toRootBook(), err
}

func (p *BookService) GetById(id string) (*root.Book, error) {
	model := bookModel{}
	err := p.collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&model)
	return model.toRootBook(), err
}
