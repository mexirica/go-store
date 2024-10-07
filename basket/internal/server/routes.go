package server

import (
	"github.com/sirupsen/logrus"
	"go-store/basket/internal/types"
	"go-store/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Logger *logrus.Logger

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(logging.LoggingMiddleware(Logger))
	r.Use(logging.ErrorMiddleware(Logger))

	r.Group("/admin").GET("/health", s.healthHandler)

	r.GET("/basket/:id", s.GetBasketHandler)
	r.POST("/basket", s.StoreBasketHandler)
	r.DELETE("/basket/:id", s.DeleteBasketHandler)

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	Logger.Info("Health endpoint called")
	result := s.db.Health()
	Logger.WithField("health_result", result).Info("Health endpoint response")
	c.JSON(http.StatusOK, result)
}

func (s *Server) GetBasketHandler(c *gin.Context) {
	id := c.Param("id")
	Logger.WithField("basket_id", id).Info("Request to get shopping basket")

	basket, err := s.repository.GetBasket(id, c)
	if err != nil {
		Logger.WithError(err).WithField("basket_id", id).Error("Error getting shopping basket")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.WithFields(logrus.Fields{
		"basket_id": id,
		"items":     len(basket.Items),
	}).Info("Shopping basket retrieved successfully")
	c.JSON(http.StatusOK, basket)
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	Logger.Info("Hello World endpoint called")
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func (s *Server) StoreBasketHandler(c *gin.Context) {
	var basket types.ShoppingCart
	if err := c.BindJSON(&basket); err != nil {
		Logger.WithError(err).Error("Error deserializing shopping basket data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Logger.WithFields(logrus.Fields{
		"userName":   basket.UserName,
		"items":      len(basket.Items),
		"totalPrice": basket.TotalPrice(),
	}).Info("Attempting to store shopping basket")

	result, err := s.repository.StoreBasket(&basket, c)
	if err != nil {
		Logger.WithError(err).WithField("userName", basket.UserName).Error("Error storing shopping basket")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.WithFields(logrus.Fields{
		"userName": basket.UserName,
		"result":   result,
	}).Info("Shopping basket stored successfully")
	c.JSON(http.StatusCreated, result)
}

func (s *Server) DeleteBasketHandler(c *gin.Context) {
	id := c.Param("id")
	Logger.WithField("basket_id", id).Info("Request to delete shopping basket")

	err := s.repository.DeleteBasket(id, c)
	if err != nil {
		Logger.WithError(err).WithField("basket_id", id).Error("Error deleting shopping basket")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.WithField("basket_id", id).Info("Shopping basket deleted successfully")
	c.JSON(http.StatusNoContent, nil)
}
