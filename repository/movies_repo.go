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

func AddMovie(newMovie entity.MoviesDTO) (int64, error) {
	db, err := config.MySQLConnection()
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO movies(title, overview, poster) VALUES (?, ?, ?)", newMovie.Title, newMovie.Overview, newMovie.Poster)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetMovieByID(id int64) (*entity.Movies, error) {
	db, err := config.MySQLConnection()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM movies WHERE id = ?", id)

	var dto entity.MoviesDTO
	errScan := row.Scan(&dto.ID, &dto.Title, &dto.Overview, &dto.Poster)
	if errScan != nil {
		return nil, errScan
	}

	movie := mapper.MapToEntity([]entity.MoviesDTO{dto})[0]
	return movie, nil
}
