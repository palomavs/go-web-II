package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/palomavs/go-web-II/cmd/server/handler"
	"github.com/palomavs/go-web-II/internal/products"
	"github.com/palomavs/go-web-II/pkg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.New(store.FileType, "./products.json")
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	pc := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	{
		pr.GET("/", pc.ValidateToken, pc.GetAll())
		pr.POST("/", pc.ValidateToken, pc.Store())
		pr.PUT("/:id", pc.ValidateToken, pc.Update())
		pr.DELETE("/:id", pc.ValidateToken, pc.Delete(false))
		pr.DELETE("/hardDelete/:id", pc.ValidateToken, pc.Delete(true))
		pr.PATCH("/:id", pc.ValidateToken, pc.UpdateNameAndPrice())
	}
	r.Run()
}
