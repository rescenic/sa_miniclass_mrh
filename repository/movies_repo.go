package repository

import (
	"sa_miniclass_mrh/config"
	"sa_miniclass_mrh/entity"
	"sa_miniclass_mrh/repository/mapper"
)

func GetDataMovies() ([]*entity.Movies, error) {
	db, err := config.MySQLConnection()
	if err != nil {
		return nil, err
	}

	rows, errRows := db.Query("SELECT * FROM movies")
	if errRows != nil {
		return nil, errRows
	}

	var result []entity.MoviesDTO
	for rows.Next() {
		var each = entity.MoviesDTO{}
		errScan := rows.Scan(&each.ID, &each.Title, &each.Overview, &each.Poster)
		if errScan != nil {
			return nil, errScan
		}

		result = append(result, each)
	}

	return mapper.MapToEntity(result), nil
}
