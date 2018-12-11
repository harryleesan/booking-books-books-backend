package server

import (
	"booking-books/books-backend/pkg"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router *mux.Router
}

func NewServer(b root.BookService) *Server {
	s := Server{router: mux.NewRouter()}
	NewBookRouter(b, s.NewSubrouter("/book"))
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) NewSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
