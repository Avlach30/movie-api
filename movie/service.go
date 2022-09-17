package movie

import "errors"

type Service interface {
	FetchAllMovieWithTags() ([]Movie, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FetchAllMovieWithTags() ([]Movie, error) {
	movies, err := service.repository.GetMoviesWithTags()
	if err != nil {
		return movies, errors.New("failed to get all movies with tags")
	}

	return movies, nil
}