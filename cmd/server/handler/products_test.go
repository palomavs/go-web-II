package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/palomavs/go-web-II/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRequest struct {
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
	Active       bool    `json:"active"`
}

type productServiceMock struct {
	mock.Mock
}

func (s *productServiceMock) GetAll(ctx context.Context) ([]domain.Product, error) {
	args := s.Called(ctx)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (s *productServiceMock) Store(ctx context.Context, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error) {
	args := s.Called(ctx, name, color, price, stock, code, published, creationDate, active)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (s *productServiceMock) Update(ctx context.Context, id int, name, color string, price float64, stock int, code string, published bool, creationDate string, active bool) (domain.Product, error) {
	args := s.Called(ctx, id, name, color, price, stock, code, published, creationDate, active)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (s *productServiceMock) UpdateNameAndPrice(ctx context.Context, id int, name string, price float64) (domain.Product, error) {
	args := s.Called(ctx, id, name, price)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (s *productServiceMock) HardDelete(ctx context.Context, id int) ([]domain.Product, error) {
	args := s.Called(ctx, id)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (s *productServiceMock) Delete(ctx context.Context, id int) ([]domain.Product, error) {
	args := s.Called(ctx, id)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func StartServer(handler *Product) *gin.Engine {
	r := gin.Default()
	pr := r.Group("/products")
	{
		pr.GET("/", handler.ValidateToken, handler.GetAll())
		pr.POST("/", handler.ValidateToken, handler.Store())
		pr.PUT("/:id", handler.ValidateToken, handler.Update())
		pr.DELETE("/:id", handler.ValidateToken, handler.Delete(false))
		pr.DELETE("/hardDelete/:id", handler.ValidateToken, handler.Delete(true))
		pr.PATCH("/:id", handler.ValidateToken, handler.UpdateNameAndPrice())
	}
	return r
}

func createRequestTest(method, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	res := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")

	return req, res
}

func TestUpdate_OK(t *testing.T) {
	serviceMock := new(productServiceMock)

	newprod := domain.Product{Id: 1, Name: "prod-1", Color: "celeste", Price: 852.33, Stock: 100, Code: "AAA", Published: true, CreationDate: "3-5-2005", Active: true}
	serviceMock.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(newprod, nil)
	productHandler := NewProduct(serviceMock)
	router := StartServer(productHandler)

	body, _ := json.Marshal(newprod)
	req, rr := createRequestTest(http.MethodPut, "/products/1", body)
	router.ServeHTTP(rr, req)

	type resp struct {
		Data domain.Product `json:"data"`
	}

	res := new(resp)
	err := json.Unmarshal(rr.Body.Bytes(), res)

	assert.NoError(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, newprod, res.Data)
}

func TestDelete_OK(t *testing.T) {
	serviceMock := new(productServiceMock)
	products := []domain.Product{{Id: 1, Name: "prod-1", Color: "celeste", Price: 852.33, Stock: 100, Code: "AAA", Published: true, CreationDate: "3-5-2005", Active: false}}
	serviceMock.On("Delete", mock.Anything, mock.Anything).Return(products, nil)
	productHandler := NewProduct(serviceMock)
	router := StartServer(productHandler)

	req, rr := createRequestTest(http.MethodDelete, "/products/1", nil)
	router.ServeHTTP(rr, req)

	type resp struct {
		Data []domain.Product `json:"data"`
	}

	res := new(resp)
	err := json.Unmarshal(rr.Body.Bytes(), res)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, products, res.Data)
}
