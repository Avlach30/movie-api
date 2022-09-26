package moviestudio

import "errors"

type Service interface {
	SaveNewStudio(input CreateNewStudioInput) (Studio, error)
	FindStudioByNumber(studioNumber int) (Studio, error)
	FindStudioByID(studioId int) (Studio, error)
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

func (service *service) FindStudioByNumber(studioNumber int) (Studio, error) {
	
	movieStudio, err := service.repository.FindStudioByNumber(studioNumber)
	if (err != nil) {
		return movieStudio, errors.New("failed to find movie studio by number")
	}

	return movieStudio, nil
}

func (service *service) FindStudioByID(studioId int) (Studio, error) {
	
	movieStudio, err := service.repository.FindStudioById(studioId)
	if (err != nil) {
		return movieStudio, errors.New("failed to find movie studio by id")
	}

	return movieStudio, nil
}