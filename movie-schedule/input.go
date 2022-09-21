package movieschedule

type CreateNewScheduleInput struct {
	MovieId   int    `json:"movie_id" binding:"required,gte=0"`
	StudioId  int    `json:"studio_id" binding:"required,gte=0"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Price     int    `json:"price" binding:"required,gte=10000"`
	Date      string `json:"date" binding:"required"`
}
