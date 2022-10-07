package movieschedule

type Service interface {
	SaveNewMovieSchedule(input CreateNewScheduleInput) (MovieSchedule, error)
	FetchPlayingNowSchedule() ([]MovieSchedule, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) SaveNewMovieSchedule(input CreateNewScheduleInput) (MovieSchedule, error) {
	
	schedule := MovieSchedule{
		MovieId: input.MovieId,
		StudioId: input.StudioId,
		StartTime: input.StartTime,
		EndTime: input.EndTime,
		Price: input.Price,
		Date: input.Date,
	}

	newSchedule, err := service.repository.SaveNewSchedule(schedule)
	if err != nil {
		return newSchedule, err
	}

	return newSchedule, nil
}

func (service *service) FetchPlayingNowSchedule() ([]MovieSchedule, error) {
	playingNowSchedules, err := service.repository.GetPlayingNowSchedule()
	if (err != nil) {
		return playingNowSchedules, err
	}

	return playingNowSchedules, nil
}