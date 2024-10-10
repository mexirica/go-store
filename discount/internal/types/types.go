package types

type Coupon struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Amount      int32  `json:"amount"`
}
