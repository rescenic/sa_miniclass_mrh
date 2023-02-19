package entity

type Movies struct {
	id       int
	title    string
	overview string
	poster   string
}

type MoviesDTO struct {
	ID       int
	Title    string
	Overview string
	Poster   string
}

func CreateMovies(dto MoviesDTO) *Movies {
	movies := &Movies{
		id:       dto.ID,
		title:    dto.Title,
		overview: dto.Overview,
		poster:   dto.Poster,
	}

	return movies
}

func (b *Movies) GetID() int {
	return b.id
}

func (b *Movies) GetTitle() string {
	return b.title
}

func (b *Movies) GetOverview() string {
	return b.overview
}

func (b *Movies) GetPoster() string {
	return b.poster
}
