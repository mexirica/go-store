package server_test

import (
	"bytes"
	"encoding/json"
	"go-store/basket/internal/repository"
	"go-store/basket/internal/types"
	"go-store/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetBasket(id string, c *gin.Context) (*types.ShoppingCart, error) {
	args := m.Called(id, c)
	return args.Get(0).(*types.ShoppingCart), args.Error(1)
}

func (m *MockRepository) StoreBasket(basket *types.ShoppingCart, c *gin.Context) (*types.ShoppingCart, error) {
	args := m.Called(basket, c)
	return args.Get(0).(*types.ShoppingCart), args.Error(1)
}

func (m *MockRepository) DeleteBasket(id string, c *gin.Context) error {
	args := m.Called(id, c)
	return args.Error(0)
}

func TestHealthHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLogger := logrus.New()
	mockRepo := new(MockRepository)

	s := &server.Server{
		logger:     mockLogger,
		repository: mockRepo,
	}

	req, _ := http.NewRequest(http.MethodGet, "/admin/health", nil)
	w := httptest.NewRecorder()
	r := s.RegisterRoutes()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBasketHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLogger := logrus.New()
	mockRepo := new(MockRepository)

	basket := &types.ShoppingCart{
		Items: []types.ShoppingCartItem{
			{Quantity: 2, Color: "red", Price: 19.99},
		},
	}

	mockRepo.On("GetBasket", "valid_id", mock.Anything).Return(basket, nil)

	s := &server.Server{
		logger:     mockLogger,
		repository: mockRepo,
	}

	req, _ := http.NewRequest(http.MethodGet, "/basket/valid_id", nil)
	w := httptest.NewRecorder()
	r := s.RegisterRoutes()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var result types.ShoppingCart
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, 1, len(result.Items))
	mockRepo.AssertExpectations(t)
}

func TestGetBasketHandler_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLogger := logrus.New()
	mockRepo := new(MockRepository)

	mockRepo.On("GetBasket", "invalid_id", mock.Anything).Return(nil, repository.ErrBasketNotFound)

	s := &server.Server{
		logger:     mockLogger,
		repository: mockRepo,
	}

	req, _ := http.NewRequest(http.MethodGet, "/basket/invalid_id", nil)
	w := httptest.NewRecorder()
	r := s.RegisterRoutes()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestStoreBasketHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLogger := logrus.New()
	mockRepo := new(MockRepository)

	basket := &types.ShoppingCart{
		Items: []types.ShoppingCartItem{
			{Quantity: 2, Color: "red", Price: 19.99},
		},
	}

	mockRepo.On("StoreBasket", basket, mock.Anything).Return(basket, nil)

	s := &server.Server{
		logger:     mockLogger,
		repository: mockRepo,
	}

	basketData, _ := json.Marshal(basket)
	req, _ := http.NewRequest(http.MethodPost, "/basket", bytes.NewBuffer(basketData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r := s.RegisterRoutes()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestStoreBasketHandler_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLogger := logrus.New()
	mockRepo := new(MockRepository)

	basket := &types.ShoppingCart{
		Items: []types.ShoppingCartItem{
			{Quantity: 2, Color: "red", Price: 19.99},
		},
	}

	mockRepo.On("StoreBasket", basket, mock.Anything).Return(nil, repository.ErrInsertFailed)

	s := &server.Server{
		logger:     mockLogger,
		repository: mockRepo,
	}

	basketData, _ := json.Marshal(basket)
	req, _ := http.NewRequest(http.MethodPost, "/basket", bytes.NewBuffer(basketData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r := s.RegisterRoutes()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBasketHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLogger := logrus.New()
	mockRepo := new(MockRepository)

	mockRepo.On("DeleteBasket", "valid_id", mock.Anything).Return(nil)

	s := &server.Server{
		logger:     mockLogger,
		repository: mockRepo,
	}

	req, _ := http.NewRequest(http.MethodDelete, "/basket/valid_id", nil)
	w := httptest.NewRecorder()
	r := s.RegisterRoutes()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBasketHandler_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockLogger := logrus.New()
	mockRepo := new(MockRepository)

	mockRepo.On("DeleteBasket", "invalid_id", mock.Anything).Return(repository.ErrBasketNotFound)

	s := &server.Server{
		logger:     mockLogger,
		repository: mockRepo,
	}

	req, _ := http.NewRequest(http.MethodDelete, "/basket/invalid_id", nil)
	w := httptest.NewRecorder()
	r := s.RegisterRoutes()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
