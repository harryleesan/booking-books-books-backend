package server

import (
	"booking-books/books-backend/pkg"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type bookRouter struct {
	bookService root.BookService
}

func NewBookRouter(b root.BookService, router *mux.Router) *mux.Router {
	bookRouter := bookRouter{b}
	router.HandleFunc("/", bookRouter.createBookHandler).Methods("PUT")
	router.HandleFunc("/title/{title}", bookRouter.getBookHandler).Methods("GET")
	router.HandleFunc("/author/{author}", bookRouter.getAuthorHandler).Methods("GET")
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
	Json(w, http.StatusOK, book)
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
