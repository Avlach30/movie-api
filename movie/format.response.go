package movie

import (
	movietag "movie-api/movie-tag"
)

type GetEachMoviesResponse struct {
	ID       int                        `json:"id"`
	Title    string                     `json:"title"`
	Overview string                     `json:"overview"`
	Poster   string                     `json:"poster"`
	Tags     []movietag.GetTagsResponse `json:"tags"`
}

func FormatResponseGetEachMovie(movie Movie) GetEachMoviesResponse {
	format := GetEachMoviesResponse{
		ID: movie.ID,
		Title: movie.Title,
		Overview: movie.Overview,
		Poster: movie.Poster,
	}
	
	return format
}

func FormatResponseGetMovies(movies []Movie) []GetEachMoviesResponse {
	moviesWithTagFormatter := []GetEachMoviesResponse{}

	for _, eachMovie := range movies {
		eachMovieFormatter := FormatResponseGetEachMovie(eachMovie)

		eachMovieFormatter.Tags = []movietag.GetTagsResponse{}

		for _, eachTag := range eachMovie.Tags {
			eachMovieFormatter.Tags = append(eachMovieFormatter.Tags, movietag.GetTagsResponse{Name: eachTag.Name, ID: eachTag.ID})
		}

		moviesWithTagFormatter = append(moviesWithTagFormatter, eachMovieFormatter)
		
	}

	return moviesWithTagFormatter
}