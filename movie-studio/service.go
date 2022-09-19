package moviestudio

import "errors"

type Service interface {
	SaveNewStudio(input CreateNewStudioInput) (Studio, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) SaveNewStudio(input CreateNewStudioInput) (Studio, error) {
	studio := Studio{
		StudioNumber: input.StudioNumber,
		SeatCapacity: input.SeatCapacity,
	}

	newStudio, err := service.repository.SaveNewStudio(studio)
	if err != nil {
		return newStudio, errors.New("failed to save new studio")
	}

	return newStudio, nil
}