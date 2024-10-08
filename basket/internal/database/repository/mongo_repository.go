package repository

import (
	"errors"
	"go-store/basket/internal/types"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
	logger *logrus.Logger
}

func NewMongoRepository(client *mongo.Client, logger *logrus.Logger) *MongoRepository {
	dbname := os.Getenv("DB_BASKET_NAME")
	return &MongoRepository{
		client: client,
		coll:   client.Database(dbname).Collection("basket"),
		logger: logger,
	}
}

func (r *MongoRepository) GetBasket(id string, c *gin.Context) (*types.ShoppingCart, error) {
	r.logger.WithField("basket_id", id).Info("Getting basket by ID")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.WithError(err).Error("Invalid basket ID")
		return nil, errors.New("invalid basket ID")
	}
	var basket types.ShoppingCart
	if err := r.coll.FindOne(c, bson.M{"_id": oid}).Decode(&basket); err != nil {
		r.logger.WithError(err).Error("Error getting basket by ID")
		return nil, err
	}
	r.logger.WithField("basket", basket).Info("Basket retrieved successfully")
	return &basket, nil
}

func (r *MongoRepository) StoreBasket(basket *types.ShoppingCart, c *gin.Context) (*types.ShoppingCart, error) {
	r.logger.WithField("basket", basket).Info("Storing basket")
	res, err := r.coll.InsertOne(c, basket)
	if err != nil {
		r.logger.WithError(err).Error("Error storing basket")
		return nil, err
	}
	basket.ID = res.InsertedID.(primitive.ObjectID)
	r.logger.WithField("basket", basket).Info("Basket stored successfully")
	return basket, nil
}

func (r *MongoRepository) DeleteBasket(id string, c *gin.Context) error {
	r.logger.WithField("basket_id", id).Info("Deleting basket by ID")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.WithError(err).Error("Invalid basket ID")
		return errors.New("invalid basket ID")
	}
	_, err = r.coll.DeleteOne(c, bson.M{"_id": oid})
	if err != nil {
		r.logger.WithError(err).Error("Error deleting basket")
		return err
	}
	r.logger.WithField("basket_id", id).Info("Basket deleted successfully")
	return nil
}
