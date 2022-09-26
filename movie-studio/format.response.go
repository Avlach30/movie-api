package moviestudio

type GetStudio struct {
	ID           int `json:"id"`
	StudioNumber int `json:"studio_number"`
	SeatCapacity int `json:"seat_capacity"`
}

func FormatCreateNewStudioResponse(studio Studio) GetStudio {
	format := GetStudio{
		ID: studio.ID,
		StudioNumber: studio.StudioNumber,
		SeatCapacity: studio.SeatCapacity,
	}

	return format
}