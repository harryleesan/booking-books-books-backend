package server

import (
	"booking-books-books-backend/pkg"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type bookRouter struct {
	bookService root.BookService
}

func BookRouter(b root.BookService, router *mux.Router) *mux.Router {
	bookRouter := bookRouter{b}
	router.HandleFunc("/", bookRouter.createBookHandler).Methods("PUT")
	router.HandleFunc("/title/{title}", bookRouter.getBookHandler).Methods("GET")
	router.HandleFunc("/author/{author}", bookRouter.getAuthorHandler).Methods("GET")
	router.HandleFunc("/id/{id}", bookRouter.getIdHandler).Methods("GET")
	router.HandleFunc("/all", bookRouter.getAllBooksHandler).Methods("GET")
	return router
}

func (br *bookRouter) createBookHandler(w http.ResponseWriter, r *http.Request) {
	book, err := decodeBook(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = br.bookService.CreateBook(&book)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, err)
}

func (br *bookRouter) getBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	title := vars["title"]
	book, err := br.bookService.GetByTitle(title)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}
	getBooks.WithLabelValues("200", "GET", "title").Inc()
	Json(w, http.StatusOK, book)
}

func (br *bookRouter) getAuthorHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	author := vars["author"]
	book, err := br.bookService.GetByAuthor(author)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}
	getBooks.WithLabelValues("200", "GET", "author").Inc()
	Json(w, http.StatusOK, book)
}

func (br *bookRouter) getIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	id := vars["id"]
	book, err := br.bookService.GetById(id)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}
	getBooks.WithLabelValues("200", "GET", "id").Inc()
	Json(w, http.StatusOK, book)
}

func (br *bookRouter) getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := br.bookService.GetAll()
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}
	getBooks.WithLabelValues("200", "GET", "all").Inc()
	Json(w, http.StatusOK, books)
}

func decodeBook(r *http.Request) (root.Book, error) {
	var b root.Book
	if r.Body == nil {
		return b, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	return b, err
}
