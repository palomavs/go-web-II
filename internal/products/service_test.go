package products

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/palomavs/go-web-II/internal/domain"
	"github.com/palomavs/go-web-II/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
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
	service := NewService(repository)

	result, err := service.GetAll(context.Background())

	assert.Equal(t, input, result, "deberían ser iguales")
	assert.Nil(t, err, "no debería dar error")
}

func TestServiceUpdate(t *testing.T) {
	prod := domain.Product{Id: 1, Name: "before change", Color: "azul", Price: 1, Stock: 1, Code: "1", Published: false, CreationDate: "1", Active: false}
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
	service := NewService(repository)

	id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive := 1, "After Update", "celeste", 2.0, 2, "2", true, "2", true
	expectedResult := domain.Product{Id: id, Name: newName, Color: newColor, Price: newPrice, Stock: newStock, Code: newCode, Published: newPublished, CreationDate: newDate, Active: newActive}

	result, errResult := service.Update(context.Background(), id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Nil(t, errResult, "no debería dar error")
	assert.True(t, storeMock.Mock.ReadCalled)
}

func TestServiceDelete(t *testing.T) {
	prod := domain.Product{Id: 1, Name: "prod1", Color: "celeste", Price: 44.44, Stock: 222, Code: "KJS4", Published: true, CreationDate: "13-12-2021", Active: true}
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
	service := NewService(repository)

	id := 1
	deletedProduct := prod
	deletedProduct.Active = false
	expectedResult := []domain.Product{deletedProduct}

	result, errResult := service.Delete(context.Background(), id)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Nil(t, errResult, "no debería dar error")
}

func TestServiceDeleteError(t *testing.T) {
	prod := domain.Product{Id: 1, Name: "prod1", Color: "celeste", Price: 44.44, Stock: 222, Code: "KJS4", Published: true, CreationDate: "13-12-2021", Active: true}
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
	service := NewService(repository)

	id := 2
	expectedResult := []domain.Product{}
	expectedError := fmt.Errorf("producto de id %d no encontrado", id)

	result, errResult := service.Delete(context.Background(), id)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.NotNil(t, errResult, "debería dar error")
}
