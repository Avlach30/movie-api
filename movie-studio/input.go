package moviestudio

type CreateNewStudioInput struct {
	StudioNumber int `json:"studio_number" binding:"required,gte=0"`
	SeatCapacity int `json:"seat_capacity" binding:"required,gte=10"`
}