package movietag

type Service interface{
	FetchAllTags() ([]Tag, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FetchAllTags() ([]Tag, error) {
	tags, err := service.repository.FindAllTags()
	if err != nil {
		return tags, err
	}

	return tags, nil
}