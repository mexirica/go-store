package repository_test

import (
	"context"
	"errors"
	"testing"

	"go-store/basket/internal/types"
	"go-store/basket/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

type MockSingleResult struct {
	mock.Mock
}

func (m *MockSingleResult) Decode(v interface{}) error {
	args := m.Called(v)
	return args.Error(0)
}

func TestStoreBasket_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	mockColl := new(MockCollection)
	mockLogger := logrus.New()

	repo := &repository.MongoRepository{
		coll:   mockColl,
		logger: mockLogger,
	}

	basket := &types.ShoppingCart{
		Items: []types.ShoppingCartItem{
			{Quantity: 2, Color: "red", Price: 19.99},
		},
	}

	mockInsertResult := &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}
	mockColl.On("InsertOne", c, basket).Return(mockInsertResult, nil)

	result, err := repo.StoreBasket(basket, c)

	assert.NoError(t, err)
	assert.Equal(t, mockInsertResult.InsertedID, result.ID)
	mockColl.AssertExpectations(t)
}

func TestStoreBasket_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	mockColl := new(MockCollection)
	mockLogger := logrus.New()

	repo := &repository.MongoRepository{
		coll:   mockColl,
		logger: mockLogger,
	}

	basket := &types.ShoppingCart{
		Items: []types.ShoppingCartItem{
			{Quantity: 2, Color: "red", Price: 19.99},
		},
	}

	mockColl.On("InsertOne", c, basket).Return(nil, errors.New("insertion error"))

	result, err := repo.StoreBasket(basket, c)

	assert.Nil(t, result)
	assert.Error(t, err)
	mockColl.AssertExpectations(t)
}

func TestGetBasket_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	mockColl := new(MockCollection)
	mockLogger := logrus.New()

	repo := &repository.MongoRepository{
		coll:   mockColl,
		logger: mockLogger,
	}

	id := primitive.NewObjectID().Hex()
	oid, _ := primitive.ObjectIDFromHex(id)

	mockSingleResult := new(MockSingleResult)
	mockColl.On("FindOne", c, bson.M{"_id": oid}).Return(mockSingleResult)

	basket := &types.ShoppingCart{}
	mockSingleResult.On("Decode", basket).Return(nil)

	result, err := repo.GetBasket(id, c)

	assert.NoError(t, err)
	assert.Equal(t, basket, result)
	mockColl.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

func TestGetBasket_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	mockLogger := logrus.New()

	repo := &repository.MongoRepository{
		logger: mockLogger,
	}

	_, err := repo.GetBasket("invalid_id", c)

	assert.Error(t, err)
	assert.Equal(t, "invalid basket ID", err.Error())
}

func TestDeleteBasket_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	mockColl := new(MockCollection)
	mockLogger := logrus.New()

	repo := &repository.MongoRepository{
		coll:   mockColl,
		logger: mockLogger,
	}

	id := primitive.NewObjectID().Hex()
	oid, _ := primitive.ObjectIDFromHex(id)

	mockColl.On("DeleteOne", c, bson.M{"_id": oid}).Return(&mongo.DeleteResult{DeletedCount: 1}, nil)

	err := repo.DeleteBasket(id, c)

	assert.NoError(t, err)
	mockColl.AssertExpectations(t)
}

func TestDeleteBasket_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c := &gin.Context{}
	mockLogger := logrus.New()

	repo := &repository.MongoRepository{
		logger: mockLogger,
	}

	err := repo.DeleteBasket("invalid_id", c)

	assert.Error(t, err)
	assert.Equal(t, "invalid basket ID", err.Error())
}
