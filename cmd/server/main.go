package main

import (
	"github.com/gin-gonic/gin"
	"github.com/palomavs/go-web-II/cmd/server/handler"
	"github.com/palomavs/go-web-II/internal/products"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)
	productsController := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.GET("/", productsController.ValidateToken, productsController.GetAll())
	pr.POST("/", productsController.ValidateToken, productsController.Store())
	r.Run()
}