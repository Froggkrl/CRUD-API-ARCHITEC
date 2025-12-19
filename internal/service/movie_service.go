package service

import (
	"CRUD-API-ARCHITEC/internal/repository"
	"CRUD-API-ARCHITEC/model"
)

type MovieService interface {
	GetAllMovies() ([]model.Movie, error)
	GetMovieBuID(id string) (*model.Movie, error)
	CreateMovie(movie *model.Movie) (*model.Movie, error)
	UpdateMovie(id string, movie *model.Movie) (*model.Movie, error)
	DeleteMovie(id string) error
}

type movieService struct {
	repo  repository.MovieRepository
	idGen *utils.IDGenerator
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{
		repo:  repo,
		idGen: utils.NewIDGenerator(),
	}
}
