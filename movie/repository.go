package movie

import (
	"errors"
	movietag "movie-api/movie-tag"
	"gorm.io/gorm"
)

type Repository interface {
	GetMoviesWithTags() ([]Movie, error)
	SaveNewMovieWithTags(movie Movie, movieTags []string) (Movie, []string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repository *repository) GetMoviesWithTags() ([]Movie, error) {
	var movies []Movie

	err := repository.db.Preload("Tags").Find(&movies).Error
	if (err != nil) {
		return movies, errors.New("failed to get all movies with tags")
	}

	return movies, nil
}

func (repository *repository) SaveNewMovieWithTags(movie Movie, movieTags []string) (Movie, []string, error) {

	transaction := repository.db.Begin()

	err := repository.db.Create(&movie).Error
	if err != nil {
		return movie, movieTags, errors.New("failed to save new movie to database")
	}
 

	for _, newTag := range movieTags {
		tag := movietag.Tag{
			Name: newTag,
		}

		err = repository.db.Create(&tag).Error
		if err != nil {
			transaction.Rollback()
			return movie, movieTags, errors.New("failed to save new tags to database")
		}

		movieTag := movietag.MovieTag{
			MovieId: movie.ID,
			TagId: tag.ID,
		}

		err = repository.db.Create(&movieTag).Error
		if err != nil {
			transaction.Rollback()
			return movie, movieTags, errors.New("failed to save new tags to database")
		}
	}

	transaction.Commit()

	return movie, movieTags, nil
}
