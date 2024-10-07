package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Price       float32            `bson:"price"`
	Category    []string           `bson:"category"`
	Description string             `bson:"description"`
	ImageFile   string             `bson:"image"`
}

type UpdateProductParams struct {
	Price       float32  `bson:"price"`
	ImageFile   string   `bson:"image"`
	Description string   `bson:"description"`
	Category    []string `bson:"category"`
}
