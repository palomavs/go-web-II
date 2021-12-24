package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/palomavs/go-web-II/internal/products"
)

const TOKEN string = "1A2B3C45D6"

type request struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Color        string  `json:"color" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
	Code         string  `json:"code" binding:"required"`
	Published    bool    `json:"published" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, products)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		newProduct, err := c.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, newProduct)
	}
}

func (c *Product) ValidateToken(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != TOKEN {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}
}
