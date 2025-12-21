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

func (s *movieService) GetAllMovie() ([]model.Movie, error) {
	return s.repo.FindAll()
}

func (s *movieService) GetMovieByID(id string) (model.Movie, error) {
	return s.repo.FindByID(id)
}

func (s *movieService) CreateMovie(movie *model.Movie) (*model.Movie, error) {
	movie.ID = s.idGen.Generate()
	if movie.Director == nil {
		movie.Director = &model.Director{
			FirstName: "",
			LastName:  "",
		}
	}
	if err := s.repo.Create(movie); err != nil {
		return nil, err
	}
	return movie, nil
}
func (s *movieService) UpdateMovie(id string, movie *model.Movie) (*model.Movie, error) {
	existingMovie, err := s.repo.FindByID(id)
	if err != nil {

		return nil, err
	}

	movie.ID = id

	if movie.Director == nil {
		movie.Director = existingMovie.Director
	}

	if err := s.repo.Update(id, movie); err != nil {
		return nil, err
	}
	return movie, nil
}

func (s *movieService) DeleteMovie(id string) error {
	return s.repo.Delete(id)
}
