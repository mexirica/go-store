package types

import (
    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type ShoppingCart struct {
    ID primitive.ObjectID `bson:"_id,omitempty"`
    UserName string `bson:"user_name"`
    Items    []ShoppingCartItem `bson:"items"`
}

type ShoppingCartItem struct {
    Quantity int `bson:"quantity"`
    Color   string `bson:"color"`
    Price float32 `bson:"price"`
    ProductName string `bson:"product_name"`
}

func (sc *ShoppingCart) TotalPrice() float32 {
    var totalPrice float32
    for _, item := range sc.Items {
        totalPrice += item.Price * float32(item.Quantity)
    }
    return totalPrice
}

type CheckoutBasket struct {
    UserName string `bson:"user_name"`
    CustomerId uuid.UUID `bson:"customer_id"`
    TotalPrice float32 `bson:"total_price"`
    ShippingAddress string `bson:"shipping_address"`
    BillingAddress string `bson:"billing_address"`
}

type ShippingAddress struct {
    FirstName string `bson:"first_name"`
    LastName string `bson:"last_name"`
    Email string `bson:"email"`
    AddressLine string `bson:"address_line"`
    City string `bson:"city"`
    State string `bson:"state"`
    ZipCode string `bson:"zip_code"`
}

type BillingAddress struct {
    CardName string `bson:"card_name"`
    CardNumber string `bson:"card_number"`
    ExpirationDate string `bson:"expiration_date"`
    CVV string `bson:"cvv"`
    PaymentMethod uint8 `bson:"payment_method"`
}