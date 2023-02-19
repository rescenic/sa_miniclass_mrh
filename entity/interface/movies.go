package _interface

type MoviesInterface interface {
	GetID() int
	GetTitle() string
	GetOverview() string
	GetPoster() string
}
