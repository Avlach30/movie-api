package movie

import (
	"errors"
)

type Service interface {
	FetchAllMovieWithTags() ([]Movie, error)
	SaveNewMovieWithTags(input CreateNewMovieInput) (Movie, error)
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

func (service *service) SaveNewMovieWithTags(input CreateNewMovieInput) (Movie, error) {
	movie := Movie{
		Title: input.Title,
		Overview: input.Overview,
		Poster: "/" + input.Poster,
		PlayUntil: input.PlayUntil,
	}

	tags := input.Tags
	
	newMovie, tags, err := service.repository.SaveNewMovieWithTags(movie, tags)
	if (err != nil || len(tags) == 0) {
		return newMovie, errors.New("failed create new movie with tags")
	}

	return newMovie, nil
}