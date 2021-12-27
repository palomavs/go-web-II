package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/palomavs/go-web-II/cmd/server/handler"
	"github.com/palomavs/go-web-II/internal/products"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	repository := products.NewRepository()
	service := products.NewService(repository)
	productsController := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	{
		pr.GET("/", productsController.ValidateToken, productsController.GetAll())
		pr.POST("/", productsController.ValidateToken, productsController.Store())
		pr.PUT("/:id", productsController.ValidateToken, productsController.Update())
		pr.DELETE("/:id", productsController.ValidateToken, productsController.Delete(false))
		pr.DELETE("/hardDelete/:id", productsController.ValidateToken, productsController.Delete(true))
		pr.PATCH("/:id", productsController.ValidateToken, productsController.UpdateNameAndPrice())
	}
	r.Run()
}
