package main

import (
	"fmt"
	"log"
	"net/http"
	"sa_miniclass_mrh/handler"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// GET request handler
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Rescenic Server API: Movies DB")
	}).Methods(http.MethodGet)

	// GET request handler for fetching movies
	route.HandleFunc("/movies/", handler.MoviesHandler).Methods(http.MethodGet)

	// POST request handler for adding new movies
	route.HandleFunc("/movies/", handler.AddMovieHandler).Methods(http.MethodPost)

	fmt.Println("http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", route))
}
