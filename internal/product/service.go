package product

import (
	"errors"

	"github.com/LucasRosello/poc-dell/internal/domain"
)

type service struct {
	repository Repository
}

type Service interface {
	GetAll() ([]domain.Product, error)
	Get(product_code string) (domain.Product, error)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() ([]domain.Product, error) {
	products, err := s.repository.GetAll()

	if err != nil {
		return []domain.Product{}, errors.New("internal Server Error")
	}

	if len(products) == 0 {
		return []domain.Product{}, errors.New("no products in the database")
	}

	return products, nil
}

func (s *service) Get(product_code string) (domain.Product, error) {
	product, err := s.repository.Get(product_code)

	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil

}
