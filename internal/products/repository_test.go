package products

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/palomavs/go-web-II/internal/domain"
	"github.com/palomavs/go-web-II/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	prod1 := domain.Product{Id: 1, Name: "prod1", Color: "celeste", Price: 44.44, Stock: 222, Code: "KJS4", Published: true, CreationDate: "13-12-2021", Active: true}
	prod2 := domain.Product{Id: 2, Name: "prod2", Color: "azul", Price: 14.14, Stock: 672, Code: "7UF4", Published: false, CreationDate: "13-12-2021", Active: true}
	input := []domain.Product{prod1, prod2}

	dataJson, _ := json.Marshal(input)
	dbStub := store.Mock{
		Data:       dataJson,
		Err:        nil,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	result, errResult := repository.GetAll(context.Background())
	assert.Equal(t, input, result, "deben ser iguales")
	assert.Nil(t, errResult, "no debería dar error")
	assert.True(t, storeMock.Mock.ReadCalled)
}

func TestUpdateNameAndPrice(t *testing.T) {
	prod := domain.Product{Id: 1, Name: "Before Change", Color: "azul", Price: 14.14, Stock: 672, Code: "7UF4", Published: false, CreationDate: "13-12-2021", Active: true}
	input := []domain.Product{prod}

	dataJson, _ := json.Marshal(input)
	dbStub := store.Mock{
		Data:       dataJson,
		Err:        nil,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	id, newName, newPrice := 1, "After Update", 100.10
	expectedResult := prod
	expectedResult.Name = newName
	expectedResult.Price = newPrice

	result, errResult := repository.UpdateNameAndPrice(context.Background(), id, newName, newPrice)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Nil(t, errResult, "no debería dar error")
	assert.True(t, storeMock.Mock.ReadCalled)

	// Validación extra de campos cambiados
	assert.Equal(t, id, result.Id, "deben ser iguales")
	assert.Equal(t, newName, result.Name, "deben ser iguales")
	assert.Equal(t, newPrice, result.Price, "deben ser iguales")
}
