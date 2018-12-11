package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AllBooksEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", AllBooksEndPoint).Methods("GET")
	// r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	// r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	// r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	// r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
