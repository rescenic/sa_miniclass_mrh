package handler

import (
	"encoding/json"
	"net/http"
	"sa_miniclass_mrh/entity"
	"sa_miniclass_mrh/repository"
	"sa_miniclass_mrh/response"
)

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getMoviesHandler(w, r)
	} else if r.Method == http.MethodPost {
		AddMovieHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Handler function for handling GET requests to /movies/ endpoint
func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := repository.GetDataMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataJson, errJson := response.MapResponseListMovies(movies)
	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(dataJson)
}

// Handler function for handling POST requests to /movies/ endpoint
func AddMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var newMovie entity.Movies
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert newMovie to an entity.MoviesDTO instance
	newMovieDTO := entity.MoviesDTO{
		ID:       newMovie.ID,
		Title:    newMovie.Title,
		Overview: newMovie.Overview,
		Poster:   newMovie.Poster,
	}

	// Add the new movie to the database
	err = repository.AddMovie(newMovieDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the added movie as JSON response
	jsonResp, err := response.MapResponseSingleMovie(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
