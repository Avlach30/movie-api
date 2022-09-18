package movie

type CreateNewMovieInput struct {
	Title     string `form:"title" binding:"required"`
	Overview  string `form:"overview" binding:"required"`
	Poster    string
	PlayUntil string `form:"play_until" binding:"required"`
	Tags      []string `form:"tags"`
}
