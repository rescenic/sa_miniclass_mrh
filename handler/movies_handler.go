package handler

import (
	"net/http"
	"sa_miniclass_mrh/repository"
	"sa_miniclass_mrh/response"
)

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := repository.GetDataMovies()
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	dataJson, errJson := response.MapResponseListMovies(movies)
	if errJson != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(errJson.Error()))
		return
	}

	_, _ = w.Write(dataJson)
}
