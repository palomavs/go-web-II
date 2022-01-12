package products

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/palomavs/go-web-II/internal/domain"
	"github.com/palomavs/go-web-II/pkg/store"
	"github.com/stretchr/testify/assert"
)

const (
	errorGetAll             = "error for GetAll"
	errorStore              = "error for Store"
	errorUpdate             = "error for Update"
	errorUpdateNameAndPrice = "error for UpdateNameAndPrice"
	errorHardDelete         = "error for HardDelete"
	errorDelete             = "error for Delete"
	errorNotFound           = "producto de id 2 no encontrado"
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

func TestGetAllError(t *testing.T) {
	expectedError := errors.New(errorGetAll)
	expectedResult := []domain.Product{}

	dbStub := store.Mock{
		Data:       nil,
		Err:        expectedError,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	result, errResult := repository.GetAll(context.Background())
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.True(t, storeMock.Mock.ReadCalled)
}

func TestStoreError(t *testing.T) {
	expectedResult := domain.Product{}
	expectedError := errors.New(errorStore)

	dbStub := store.Mock{
		Data:       nil,
		Err:        expectedError,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive := 1, "prod1", "celeste", 44.40, 222, "K4KH", true, "22-01-22", true

	result, errResult := repository.Store(context.Background(), id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.NotNil(t, errResult, "debería dar error")
}

func TestUpdateError(t *testing.T) {
	expectedResult := domain.Product{}
	expectedError := errors.New(errorUpdate)

	dbStub := store.Mock{
		Data:       nil,
		Err:        expectedError,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive := 1, "After Update", "celeste", 2.0, 2, "2", true, "2", true

	result, errResult := repository.Update(context.Background(), id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive)
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.Equal(t, expectedResult, result, "deben ser iguales")
}

func TestUpdateNotFound(t *testing.T) {
	expectedResult := domain.Product{}
	expectedError := errors.New(errorNotFound)
	input := []domain.Product{}

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

	id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive := 2, "After Update", "celeste", 2.0, 2, "2", true, "2", true

	result, errResult := repository.Update(context.Background(), id, newName, newColor, newPrice, newStock, newCode, newPublished, newDate, newActive)
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.Equal(t, expectedResult, result, "deben ser iguales")
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

func TestUpdateNameAndPriceError(t *testing.T) {
	expectedError := errors.New(errorUpdateNameAndPrice)
	expectedResult := domain.Product{}

	dbStub := store.Mock{
		Data:       nil,
		Err:        expectedError,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	id, newName, newPrice := 1, "After Update", 100.10

	result, errResult := repository.UpdateNameAndPrice(context.Background(), id, newName, newPrice)
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.True(t, storeMock.Mock.ReadCalled)
}

func TestUpdateNameAndPriceNotFound(t *testing.T) {
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

	id, newName, newPrice := 2, "After Update", 100.10
	expectedError := errors.New(errorNotFound)
	expectedResult := domain.Product{}

	result, errResult := repository.UpdateNameAndPrice(context.Background(), id, newName, newPrice)
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.True(t, storeMock.Mock.ReadCalled)
}

func TestHardDeleteError(t *testing.T) {
	expectedError := errors.New(errorHardDelete)

	dbStub := store.Mock{
		Data:       nil,
		Err:        expectedError,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	id := 1
	expectedResult := []domain.Product{}

	result, errResult := repository.HardDelete(context.Background(), id)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.NotNil(t, errResult, "debería dar error")
}

func TestDeleteError(t *testing.T) {
	expectedError := errors.New(errorDelete)

	dbStub := store.Mock{
		Data:       nil,
		Err:        expectedError,
		ReadCalled: false,
	}
	storeMock := store.FileStore{
		FileName: "",
		Mock:     &dbStub,
	}
	repository := NewRepository(&storeMock)

	id := 1
	expectedResult := []domain.Product{}

	result, errResult := repository.Delete(context.Background(), id)
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Equal(t, expectedError, errResult, "deben ser iguales")
	assert.NotNil(t, errResult, "debería dar error")
}

func TestLastID(t *testing.T) {
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
	expectedResult := 2

	result, errResult := repository.LastID(context.Background())
	assert.Equal(t, expectedResult, result, "deben ser iguales")
	assert.Nil(t, errResult, "no debe dar error")
}
