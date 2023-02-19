package entity

type Movies struct {
	ID       int
	Title    string
	Overview string
	Poster   string
}

type MoviesDTO struct {
	ID       int
	Title    string
	Overview string
	Poster   string
}

func CreateMovies(dto MoviesDTO) *Movies {
	movies := &Movies{
		ID:       dto.ID,
		Title:    dto.Title,
		Overview: dto.Overview,
		Poster:   dto.Poster,
	}

	return movies
}

func (m *Movies) GetID() int {
	return m.ID
}

func (m *Movies) GetTitle() string {
	return m.Title
}

func (m *Movies) GetOverview() string {
	return m.Overview
}

func (m *Movies) GetPoster() string {
	return m.Poster
}
