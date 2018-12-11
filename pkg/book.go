package root

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookService interface {
	CreateBook(b *Book) error
	GetByTitle(title string) (*Book, error)
	GetByAuthor(author string) (*Book, error)
	GetById(id string) (*Book, error)
}
