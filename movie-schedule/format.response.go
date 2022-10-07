package movieschedule

import (
	"movie-api/movie"
	moviestudio "movie-api/movie-studio"
)

type CreateNewScheduleResponse struct {
	ID        int                   `json:"id"`
	Movie     movie.GetMovie        `json:"movie"`
	Studio    moviestudio.GetStudio `json:"studio"`
	StartTime string                `json:"start_time"`
	EndTime   string                `json:"end_time"`
	Price     int                   `json:"price"`
	Date      string                `json:"date"`
}

type GetPlayingNowScheduleResponse struct {
	ID        int                   `json:"id"`
	Movie     movie.MovieResponse   `json:"movie"`
	Studio    moviestudio.GetStudio `json:"studio"`
	StartTime string                `json:"start_time"`
	EndTime   string                `json:"end_time"`
	Price     int                   `json:"price"`
	Date      string                `json:"date"`
}

func FormatCreateNewScheduleResponse(schedule MovieSchedule, movie movie.Movie, studio moviestudio.Studio) CreateNewScheduleResponse {
	format := CreateNewScheduleResponse{
		ID:        schedule.ID,
		StartTime: schedule.StartTime,
		EndTime:   schedule.EndTime,
		Price:     schedule.Price,
		Date:      schedule.Date,
	}

	format.Movie.Title = movie.Title
	format.Movie.Poster = movie.Poster

	format.Studio.ID = studio.ID
	format.Studio.StudioNumber = studio.StudioNumber
	format.Studio.SeatCapacity = studio.SeatCapacity

	return format
}

func FormatGetEachPlayingNowScheduleResponse(schedule MovieSchedule, movie movie.Movie, studio moviestudio.Studio) GetPlayingNowScheduleResponse {
	format := GetPlayingNowScheduleResponse{
		ID: schedule.ID,
		StartTime: schedule.StartTime,
		EndTime: schedule.EndTime,
		Price: schedule.Price,
		Date: schedule.Date,
	}

	format.Movie.ID = movie.ID
	format.Movie.Title = movie.Title
	format.Movie.Overview = movie.Overview
	format.Movie.Poster = movie.Poster
	format.Movie.PlayUntil = movie.PlayUntil

	format.Studio.ID = studio.ID
	format.Studio.StudioNumber = studio.StudioNumber
	format.Studio.SeatCapacity = studio.SeatCapacity

	return format
}

func FormatGetPlayingNowSchedulesResponse(schedules []MovieSchedule) []GetPlayingNowScheduleResponse {
	format := []GetPlayingNowScheduleResponse{}

	for _, schedule := range schedules {
		eachScheduleFormatter := FormatGetEachPlayingNowScheduleResponse(schedule, schedule.Movie, schedule.Studio)

		format = append(format, eachScheduleFormatter)
	}

	return format
}
