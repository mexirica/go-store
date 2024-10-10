package repository

import "go-store/discount/internal/types"

type Repository interface {
	GetDiscount(string) (*types.Coupon, error)
	CreateDiscount(*types.Coupon) (*types.Coupon, error)
	UpdateDiscount(*types.Coupon) error
	DeleteDiscount(string) error
}
