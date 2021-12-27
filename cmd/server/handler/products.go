package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palomavs/go-web-II/internal/products"
)

const TOKEN = "1A2B3C45D6"

type request struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
	Active       bool    `json:"active"`
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

		newProduct, err := c.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate, req.Active)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, newProduct)
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err = ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		{ // Validaciones - se podría hacer con el required pero choca con el patch
			if req.Name == "" {
				ctx.JSON(400, gin.H{"error": "debe proveer un nombre de producto"})
				return
			}
			if req.Color == "" {
				ctx.JSON(400, gin.H{"error": "debe proveer un color para el producto"})
				return
			}
			if req.Price == 0 {
				ctx.JSON(400, gin.H{"error": "debe proveer un precio para el producto"})
				return
			}
			if req.Stock == 0 {
				ctx.JSON(400, gin.H{"error": "debe proveer un stock para el producto"})
				return
			}
			if req.Code == "" {
				ctx.JSON(400, gin.H{"error": "debe proveer un código para el producto"})
				return
			}
			if req.CreationDate == "" {
				ctx.JSON(400, gin.H{"error": "debe proveer una fecha de creación para el producto"})
				return
			}
		}

		productUpdated, err := c.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate, req.Active)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, productUpdated)
	}
}

func (c *Product) Delete(hardDelete bool) gin.HandlerFunc {
	if hardDelete {
		return func(ctx *gin.Context) {
			id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
			if err != nil {
				ctx.JSON(400, gin.H{"error": "invalid ID"})
				return
			}

			products, err := c.service.HardDelete(int(id))
			if err != nil {
				ctx.JSON(404, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(200, products)
			//ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
		}
	} else {
		return func(ctx *gin.Context) {
			id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
			if err != nil {
				ctx.JSON(400, gin.H{"error": "invalid ID"})
				return
			}

			products, err := c.service.Delete(int(id))
			if err != nil {
				ctx.JSON(404, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(200, products)
		}
	}
}

func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" || req.Price == 0 {
			ctx.JSON(400, gin.H{"error": "debe proveer un nombre y precio válidos"})
		}

		updatedProduct, err := c.service.UpdateNameAndPrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, updatedProduct)
	}
}

func (c *Product) ValidateToken(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != TOKEN {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}
}
