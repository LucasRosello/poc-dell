package product

type service struct {
	repository Repository
}

type Service interface {
	Hello() string
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Hello() string {
	return s.repository.Hello()
}
