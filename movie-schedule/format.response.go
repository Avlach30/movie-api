package movieschedule

import (
	"movie-api/movie"
	moviestudio "movie-api/movie-studio"
)

type CreateNewScheduleResponse struct {
	ID        int                `json:"id"`
	Movie     movie.GetMovie        `json:"movie"`
	Studio    moviestudio.GetStudio `json:"studio"`
	StartTime string             `json:"start_time"`
	EndTime   string             `json:"end_time"`
	Price     int                `json:"price"`
	Date      string             `json:"date"`
}

func FormatCreateNewScheduleResponse(schedule MovieSchedule, movie movie.Movie, studio moviestudio.Studio) CreateNewScheduleResponse {
	format := CreateNewScheduleResponse{
		ID: schedule.ID,
		StartTime: schedule.StartTime,
		EndTime: schedule.EndTime,
		Price: schedule.Price,
		Date: schedule.Date,
	}

	format.Movie.Title = movie.Title
	format.Movie.Poster = movie.Poster

	format.Studio.ID = studio.ID
	format.Studio.StudioNumber = studio.StudioNumber
	format.Studio.SeatCapacity = studio.SeatCapacity

	return format
}
