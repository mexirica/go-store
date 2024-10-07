package repository

import (
	"github.com/gin-gonic/gin"
	"go-store/catalog/internal/types"
)

type Product = types.Product
type Filter = map[string]any
type UpdateParams = types.UpdateProductParams

type Repository interface {
	CreateProduct(product *Product, c *gin.Context) (*Product, error)
	GetProductById(id string, c *gin.Context) (*Product, error)
	GetProducts(c *gin.Context) ([]*Product, error)
	UpdateProduct(filter Filter, params *UpdateParams, c *gin.Context) (*Product, error)
	DeleteProduct(id string, c *gin.Context) error
	GetProductsByCategory(category string, c *gin.Context) ([]*Product, error)
}
