package mapper

import "sa_miniclass_mrh/entity"

func MapToEntity(listDTO []entity.MoviesDTO) []*entity.Movies {
	var listMovies []*entity.Movies
	for _, dto := range listDTO {
		movies := entity.CreateMovies(dto)
		listMovies = append(listMovies, movies)
	}

	return listMovies
}
