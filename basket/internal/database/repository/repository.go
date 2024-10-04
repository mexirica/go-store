package repository

import (
	"go-store/basket/internal/types"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetBasket (id string, ctx *gin.Context) (*types.ShoppingCart, error)
	StoreBasket (basket *types.ShoppingCart, ctx *gin.Context) (*types.ShoppingCart, error)
	DeleteBasket (id string, ctx *gin.Context) error
}