package repository

import (
	"CRUD-API-ARCHITEC/model"
	"errors"
	"sync"
)

var (
	ErrMovieNotFound = errors.New("movie not found")
	ErrMovieExists   = errors.New("movie already exists")
)

type MovieRepository interface {
	FindAll() ([]model.Movie, error)
	FindByID(id string) (*model.Movie, error)
	Create(model *model.Movie) error
	Update(is string, model *model.Movie) error
	Delete(id string) error
}

type MemoryRepository struct {
	movies map[string]model.Movie
	mu     sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		movies: make(map[string]model.Movie),
	}
}

func (r *MemoryRepository) InitWithSampleData() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.movies["1"] = model.Movie{
		ID:    "1",
		Title: "Iron man",
		Director: &model.Director{
			FirstName: "Jon",
			LastName:  "Favreau",
		},
	}
}

func (r *MemoryRepository) FindAll() ([]model.Movie, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	movies := make([]model.Movie, 0, len(r.movies))
	for _, movie := range r.movies {
		movies = append(movies, movie)
	}
	return movies, nil
}

func (r *MemoryRepository) FindByID(id string) (*model.Movie, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	movie, exists := r.movies[id]
	if !exists {
		return nil, ErrMovieNotFound
	}
	return &movie, nil
}

func (r *MemoryRepository) Create(movie *model.Movie) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.movies[movie.ID]; exists {
		return ErrMovieExists
	}

	r.movies[movie.ID] = *movie
	return nil
}

func (r *MemoryRepository) Update(id string, movie *model.Movie) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.movies[id]; !exists {
		return ErrMovieNotFound
	}

	r.movies[id] = *movie
	return nil
}

func (r *MemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.movies[id]; !exists {
		return ErrMovieNotFound
	}
	delete(r.movies, id)
	return nil
}
