package server

import (
	"booking-books-books-backend/pkg"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func NewServer(b root.BookService) *Server {
	s := Server{router: mux.NewRouter()}
	BookRouter(b, s.NewSubrouter("/book"))
	MetricsRouter(s.router)
	return &s
}

func (s *Server) Start() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) NewSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
