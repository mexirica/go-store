package repository

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoRepository represents a repository for managing products in MongoDB.
type MongoRepository struct {
	client *mongo.Client
	coll   *mongo.Collection
	logger *logrus.Logger
}

// NewMongoRepository creates a new instance of MongoRepository.
// It initializes the collection using the database name from the environment variable DB_CATALOG_NAME.
func NewMongoRepository(client *mongo.Client, logger *logrus.Logger) *MongoRepository {
	dbname := os.Getenv("DB_CATALOG_NAME")
	return &MongoRepository{
		client: client,
		coll:   client.Database(dbname).Collection("catalog"),
		logger: logger,
	}
}

// CreateProduct inserts a new product into the MongoDB collection.
// It returns the created product or an error if the operation fails.
func (r *MongoRepository) CreateProduct(product *Product, c *gin.Context) (*Product, error) {
	if product == nil {
		r.logger.Error("Product is nil")
		return nil, errors.New("product cannot be nil")
	}
	_, err := r.coll.InsertOne(c, product)
	if err != nil {
		r.logger.WithError(err).Error("Error creating product")
		return nil, err
	}
	return product, nil
}

// GetProductById retrieves a product by its ID from the MongoDB collection.
// It returns the product or an error if the operation fails.
func (r *MongoRepository) GetProductById(id string, c *gin.Context) (*Product, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.WithError(err).Error("Invalid product ID")
		return nil, errors.New("invalid product ID")
	}
	var product Product
	err = r.coll.FindOne(c, bson.M{"_id": objectId}).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			r.logger.WithError(err).Error("Product not found")
			return nil, errors.New("product not found")
		}
		r.logger.WithError(err).Error("Error getting product by ID")
		return nil, err
	}
	return &product, nil
}

// GetProducts retrieves all products from the MongoDB collection.
// It returns a slice of products or an error if the operation fails.
func (r *MongoRepository) GetProducts(c *gin.Context) ([]*Product, error) {
	cursor, err := r.coll.Find(c, bson.M{})
	if err != nil {
		r.logger.WithError(err).Error("Error getting products")
		return nil, err
	}
	defer cursor.Close(c)
	var products []*Product
	err = cursor.All(c, &products)
	if err != nil {
		r.logger.WithError(err).Error("Error decoding products")
		return nil, err
	}
	return products, nil
}

// UpdateProduct updates an existing product in the MongoDB collection based on the provided filter and update parameters.
// It returns the updated product or an error if the operation fails.
func (r *MongoRepository) UpdateProduct(filter Filter, params *UpdateParams, c *gin.Context) (*Product, error) {
	id, ok := filter["id"].(string)
	if !ok || id == "" {
		r.logger.Error("Invalid filter ID")
		return nil, errors.New("invalid filter ID")
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.WithError(err).Error("Invalid product ID")
		return nil, errors.New("invalid product ID")
	}
	update := bson.M{
		"$set": bson.M{
			"price":       params.Price,
			"image":       params.ImageFile,
			"description": params.Description,
			"category":    params.Category,
		},
	}
	opts := options.Update().SetUpsert(false)
	_, err = r.coll.UpdateOne(c, bson.M{"_id": objectId}, update, opts)
	if err != nil {
		r.logger.WithError(err).Error("Error updating product")
		return nil, err
	}
	return r.GetProductById(id, c)
}

// DeleteProduct deletes a product by its ID from the MongoDB collection.
// It returns an error if the operation fails.
func (r *MongoRepository) DeleteProduct(id string, c *gin.Context) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.WithError(err).Error("Invalid product ID")
		return errors.New("invalid product ID")
	}
	_, err = r.coll.DeleteOne(c, bson.M{"_id": objectId})
	if err != nil {
		r.logger.WithError(err).Error("Error deleting product")
		return err
	}
	return nil
}

// GetProductsByCategory retrieves products by their category from the MongoDB collection.
// It returns a slice of products or an error if the operation fails.
func (r *MongoRepository) GetProductsByCategory(category string, c *gin.Context) ([]*Product, error) {
	if category == "" {
		r.logger.Error("Category cannot be empty")
		return nil, errors.New("category cannot be empty")
	}
	cursor, err := r.coll.Find(c, bson.M{"category": category})
	if err != nil {
		r.logger.WithError(err).Error("Error getting products by category")
		return nil, err
	}
	defer cursor.Close(c)
	var products []*Product
	err = cursor.All(c, &products)
	if err != nil {
		r.logger.WithError(err).Error("Error decoding products by category")
		return nil, err
	}
	return products, nil
}
