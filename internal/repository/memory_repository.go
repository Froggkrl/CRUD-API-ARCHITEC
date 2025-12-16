package repository

import (
	"CRUD-API-ARCHITEC/model"
	"errors"
)

var (
	ErrMovieNotFound = errors.New("movie not found")
	ErrMovieExists   = errors.New("movie already exists")
)

type MovieRepository interface {
	FindAll() ([]model.Movie, error)
	FindByID() (*model.Movie, error)
}
