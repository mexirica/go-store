package repository

import (
	"go-store/basket/internal/types"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoBasketRepository struct {
	client *mongo.Client
	coll  *mongo.Collection
}

func NewMongoBasketRepository(client *mongo.Client) *MongoBasketRepository {
	dbname := os.Getenv("DB_BASKET_NAME")
	return &MongoBasketRepository{
		client: client,
		coll: client.Database(dbname).Collection("basket"),
	}
}

func (r *MongoBasketRepository) GetBasket(id string, c *gin.Context) (*types.ShoppingCart, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var basket types.ShoppingCart
	if err := r.coll.FindOne(c, bson.M{"_id": oid}).Decode(&basket); err != nil {
		return nil, err
	}
	return &basket, nil
}

func (r *MongoBasketRepository) StoreBasket(basket *types.ShoppingCart, c *gin.Context) (*types.ShoppingCart, error) {
	res, err := r.coll.InsertOne(c, basket)
	if err != nil {
		return nil, err
	}
	basket.ID = res.InsertedID.(primitive.ObjectID)
	return basket, nil
}

func (r *MongoBasketRepository) DeleteBasket(id string, c *gin.Context) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.DeleteOne(c, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}