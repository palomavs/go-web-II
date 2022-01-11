package products

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/palomavs/go-web-II/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	prod1        = domain.Product{Id: 1, Name: "prod1", Color: "celeste", Price: 44.44, Stock: 222, Code: "KJS4", Published: true, CreationDate: "13-12-2021", Active: true}
	prod2        = domain.Product{Id: 2, Name: "prod2", Color: "azul", Price: 14.14, Stock: 672, Code: "7UF4", Published: false, CreationDate: "13-12-2021", Active: true}
	prod3        = domain.Product{Id: 3, Name: "Before Update", Color: "azul", Price: 14.14, Stock: 672, Code: "7UF4", Published: false, CreationDate: "13-12-2021", Active: true}
	productsList = []domain.Product{prod1, prod2}
)

type stubStore struct{}
type mockStore struct {
	spyCalled bool
}

func (s *stubStore) Read(data interface{}) error {
	prods, _ := json.Marshal(productsList)
	json.Unmarshal(prods, &data)

	return nil
}

func (s *stubStore) Write(data interface{}) error {
	return nil
}

func (m *mockStore) Read(data interface{}) error {
	prods, _ := json.Marshal([]domain.Product{prod3})
	json.Unmarshal(prods, &data)
	m.spyCalled = true

	return nil
}

func (m *mockStore) Write(data interface{}) error { return nil }
func TestGetAll(t *testing.T) {
	myStubStore := &stubStore{}
	repository := NewRepository(myStubStore)

	result, errResult := repository.GetAll(context.Background())
	assert.Equal(t, productsList, result, "deben ser iguales")
	assert.Nil(t, errResult, "no debería dar error")
}

func TestUpdate(t *testing.T) {
	myMockStore := &mockStore{}
	repository := NewRepository(myMockStore)

	id, newName, newPrice := 3, "After Update", 100.10
	expectedResult := prod3
	expectedResult.Name = newName
	expectedResult.Price = newPrice

	result, errResult := repository.UpdateNameAndPrice(context.Background(), id, newName, newPrice)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Nil(t, errResult, "no debería dar error")
	assert.True(t, myMockStore.spyCalled)

	// Validación extra de campos cambiados
	assert.Equal(t, id, result.Id, "deben ser iguales")
	assert.Equal(t, newName, result.Name, "deben ser iguales")
	assert.Equal(t, newPrice, result.Price, "deben ser iguales")
}
