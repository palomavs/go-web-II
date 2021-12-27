package products

import "fmt"

var products []Product
var lastID int

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
	Active       bool    `json:"active"`
}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error)
	LastID() (int, error)
	Update(id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	HardDelete(id int) ([]Product, error)
	Delete(id int) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error) {
	newProduct := Product{id, name, color, price, stock, code, published, creationDate, active}
	products = append(products, newProduct)
	lastID++

	return newProduct, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (Product, error) {
	found := false
	updatedProduct := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreationDate: creationDate, Active: active}

	for i := range products {
		if products[i].Id == id {
			found = true

			updatedProduct.Id = id
			products[i] = updatedProduct

			break
		}
	}

	if !found {
		return Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}
	return updatedProduct, nil
}

func (r *repository) HardDelete(id int) ([]Product, error) {
	found := false
	var index int

	for i := range products {
		if products[i].Id == id {
			found = true
			index = i
			break
		}
	}

	if !found {
		return []Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}

	products = append(products[:index], products[index+1:]...)
	return products, nil
}

func (r *repository) Delete(id int) ([]Product, error) {
	found := false

	for i := range products {
		if products[i].Id == id {
			found = true
			products[i].Active = false
			break
		}
	}

	if !found {
		return []Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}

	return products, nil
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	found := false
	var index int

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
		return Product{}, fmt.Errorf("producto de id %d no encontrado", id)
	}
	return products[index], nil
}
