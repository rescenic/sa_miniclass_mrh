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
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Rescenic Server API: Movies DB")
	}).Methods(http.MethodGet)
	route.HandleFunc("/movies/", handler.MoviesHandler).Methods(http.MethodGet)

	fmt.Println("http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", route))
}
