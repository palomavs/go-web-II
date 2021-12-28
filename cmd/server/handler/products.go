package handler

import (
	"context"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palomavs/go-web-II/internal/products"
	"github.com/palomavs/go-web-II/pkg/web"
)

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

// ListProducts godoc
// @Summary Lists products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.service.GetAll(context.Background())
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, products, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /products [post]
func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		{
			if req.Name == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un nombre de producto"))
				return
			}
			if req.Color == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un color para el producto"))
				return
			}
			if req.Price == 0 {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un precio para el producto"))
				return
			}
			if req.Stock == 0 {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un stock para el producto"))
				return
			}
			if req.Code == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un código para el producto"))
				return
			}
			if req.CreationDate == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer una fecha de creación para el producto"))
				return
			}
		}

		newProduct, err := c.service.Store(context.Background(), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate, req.Active)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, newProduct, ""))
	}
}

// DeleteProducts godoc
// @Summary Removes product based on given ID
// @Tags Products
// @Description removes products
// @Produce json
// @Param token header string true "token"
// @Param id path integer true "product id to be removed"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /products/{id} [delete]
func (c *Product) Delete(hardDelete bool) gin.HandlerFunc {
	if hardDelete {
		return func(ctx *gin.Context) {
			id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
				return
			}

			products, err := c.service.HardDelete(context.Background(), int(id))
			if err != nil {
				ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
				return
			}

			ctx.JSON(200, web.NewResponse(200, products, ""))
			//ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
		}
	} else {
		return func(ctx *gin.Context) {
			id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
				return
			}

			products, err := c.service.Delete(context.Background(), int(id))
			if err != nil {
				ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
				return
			}
			ctx.JSON(200, web.NewResponse(200, products, ""))
		}
	}
}

// UpdateProducts godoc
// @Summary Updates product based on given ID
// @Tags Products
// @Description updates products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path integer true "product id to be updated"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /products/{id} [put]
func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}

		var req request
		if err = ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		{
			if req.Name == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un nombre de producto"))
				return
			}
			if req.Color == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un color para el producto"))
				return
			}
			if req.Price == 0 {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un precio para el producto"))
				return
			}
			if req.Stock == 0 {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un stock para el producto"))
				return
			}
			if req.Code == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un código para el producto"))
				return
			}
			if req.CreationDate == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "debe proveer una fecha de creación para el producto"))
				return
			}
		}

		productUpdated, err := c.service.Update(context.Background(), int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate, req.Active)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, productUpdated, ""))
	}
}

// UpdateNameAndPriceProducts godoc
// @Summary Updates name and price of product based on given ID
// @Tags Products
// @Description updates name and price of products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path integer true "product id to be updated"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Failure 404 {object} web.Response
// @Router /products/{id} [patch]
func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un nombre de producto"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "debe proveer un precio para el producto"))
			return
		}

		updatedProduct, err := c.service.UpdateNameAndPrice(context.Background(), int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, updatedProduct, ""))
	}
}

func (c *Product) ValidateToken(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "no tiene permisos para realizar la petición solicitada"))
		return
	}
}
