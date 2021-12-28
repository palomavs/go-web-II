package products

import (
	"context"

	"github.com/palomavs/go-web-II/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error)
	Update(ctx context.Context, id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error)
	UpdateNameAndPrice(ctx context.Context, id int, name string, price float64) (domain.Product, error)
	HardDelete(ctx context.Context, id int) ([]domain.Product, error)
	Delete(ctx context.Context, id int) ([]domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Store(ctx context.Context, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error) {
	lastID, err := s.repository.LastID(ctx)
	if err != nil {
		return domain.Product{}, err
	}
	lastID++

	newProduct, err := s.repository.Store(ctx, lastID, name, color, price, stock, code, published, creationDate, active)
	if err != nil {
		return domain.Product{}, err
	}

	return newProduct, nil
}

func (s *service) Update(ctx context.Context, id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error) {
	return s.repository.Update(ctx, id, name, color, price, stock, code, published, creationDate, active)
}

func (s *service) HardDelete(ctx context.Context, id int) ([]domain.Product, error) {
	return s.repository.HardDelete(ctx, id)
}

func (s *service) Delete(ctx context.Context, id int) ([]domain.Product, error) {
	return s.repository.Delete(ctx, id)
}

func (s *service) UpdateNameAndPrice(ctx context.Context, id int, name string, price float64) (domain.Product, error) {
	return s.repository.UpdateNameAndPrice(ctx, id, name, price)
}
