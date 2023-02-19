package response

import (
	"encoding/json"
	"sa_miniclass_mrh/entity"
	_interface "sa_miniclass_mrh/entity/interface"
)

type JSONListMovies struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Poster   string `json:"poster"`
}

func MapResponseListMovies(listMovies []*entity.Movies) ([]byte, error) {
	var dataJson []*JSONListMovies
	for _, movies := range listMovies {
		var jsonSingle = EntityToJson(movies)

		dataJson = append(dataJson, jsonSingle)
	}

	byteJson, err := json.Marshal(dataJson)
	if err != nil {
		return nil, err
	}

	return byteJson, nil
}

func EntityToJson(movies _interface.MoviesInterface) *JSONListMovies {
	jsonSingle := &JSONListMovies{
		ID:       movies.GetID(),
		Title:    movies.GetTitle(),
		Overview: movies.GetOverview(),
		Poster:   movies.GetPoster(),
	}

	return jsonSingle
}

func MapResponseSingleMovie(movie *entity.Movies) ([]byte, error) {
	jsonSingle := EntityToJson(movie)

	byteJson, err := json.Marshal(jsonSingle)
	if err != nil {
		return nil, err
	}

	return byteJson, nil
}
