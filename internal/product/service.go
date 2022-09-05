package product

import (
	"errors"

	"github.com/LucasRosello/poc-dell/internal/domain"
)

type service struct {
	repository Repository
}

type Service interface {
	Hello() ([]domain.Product, error)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Hello() ([]domain.Product, error) {
	products, err := s.repository.GetAll()

	if err != nil {
		return []domain.Product{}, errors.New("internal Server Error")
	}

	if len(products) == 0 {
		return []domain.Product{}, errors.New("no products in the database")
	}

	return products, nil
}
