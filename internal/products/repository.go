package products

import (
	"context"
	"fmt"

	"github.com/palomavs/go-web-II/internal/domain"
	"github.com/palomavs/go-web-II/pkg/store"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error)
	LastID(ctx context.Context) (int, error)
	Update(ctx context.Context, id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error)
	UpdateNameAndPrice(ctx context.Context, id int, name string, price float64) (domain.Product, error)
	HardDelete(ctx context.Context, id int) ([]domain.Product, error)
	Delete(ctx context.Context, id int) ([]domain.Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	if err := r.db.Read(&products); err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (r *repository) Store(ctx context.Context, id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error) {
	newProduct := domain.Product{id, name, color, price, stock, code, published, creationDate, active}
	var products []domain.Product

	//Lo vamos a sobreescribir completo, así que necesitamos leerlo antes
	err := r.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}

	//Lo escribimos
	products = append(products, newProduct)
	err = r.db.Write(products)
	if err != nil {
		return domain.Product{}, err
	}

	return newProduct, nil
}

func (r *repository) LastID(ctx context.Context) (int, error) {
	var products []domain.Product

	err := r.db.Read(&products)
	if err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 0, nil
	}

	return products[len(products)-1].Id, nil
}

func (r *repository) Update(ctx context.Context, id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error) {
	updatedProduct := domain.Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreationDate: creationDate, Active: active}
	var products []domain.Product
	found := false

	//Lo vamos a sobreescribir completo, así que necesitamos leerlo antes
	err := r.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}

	for i := range products {
		if products[i].Id == id {
			found = true
			updatedProduct.Id = id
			products[i] = updatedProduct
			break
		}
	}

	if !found {
		return domain.Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}

	err = r.db.Write(products)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (r *repository) UpdateNameAndPrice(ctx context.Context, id int, name string, price float64) (domain.Product, error) {
	found := false
	var products []domain.Product
	var index int

	//Lo vamos a sobreescribir completo, así que necesitamos leerlo antes
	err := r.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}

	for i := range products {
		if products[i].Id == id {
			found = true
			index = i
			products[i].Name = name
			products[i].Price = price
			break
		}
	}

	if !found {
		return domain.Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}

	err = r.db.Write(products)
	if err != nil {
		return domain.Product{}, err
	}

	return products[index], nil
}

func (r *repository) HardDelete(ctx context.Context, id int) ([]domain.Product, error) {
	found := false
	var index int
	var products []domain.Product

	err := r.db.Read(&products)
	if err != nil {
		return []domain.Product{}, err
	}

	for i := range products {
		if products[i].Id == id {
			found = true
			index = i
			break
		}
	}

	if !found {
		return []domain.Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}

	products = append(products[:index], products[index+1:]...)
	err = r.db.Write(products)
	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}

func (r *repository) Delete(ctx context.Context, id int) ([]domain.Product, error) {
	found := false
	var products []domain.Product

	err := r.db.Read(&products)
	if err != nil {
		return []domain.Product{}, err
	}

	for i := range products {
		if products[i].Id == id {
			found = true
			products[i].Active = false
			break
		}
	}

	if !found {
		return []domain.Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}

	err = r.db.Write(products)
	if err != nil {
		return []domain.Product{}, err
	}

	return products, nil
}
