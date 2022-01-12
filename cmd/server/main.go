package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/palomavs/go-web-II/cmd/server/handler"
	"github.com/palomavs/go-web-II/docs"
	"github.com/palomavs/go-web-II/internal/products"
	"github.com/palomavs/go-web-II/pkg/store"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Bootcamp - GO Web Module API
// @version 1.0
// @description This API handles MELI products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/Support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	{
		pr.GET("/", pc.ValidateToken, pc.GetAll())
		pr.POST("/", pc.ValidateToken, pc.Store())
		pr.PUT("/:id", pc.ValidateToken, pc.Update())
		pr.DELETE("/:id", pc.ValidateToken, pc.Delete(false))
		pr.DELETE("/hardDelete/:id", pc.ValidateToken, pc.Delete(true))
		pr.PATCH("/:id", pc.ValidateToken, pc.UpdateNameAndPrice())
	}
	err = r.Run()
	if err != nil {
		log.Fatal("error al intentar correr el server")
	}
}
