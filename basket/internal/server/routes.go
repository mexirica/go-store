package server

import (
	"go-store/basket/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Group("/admin").GET("/health", s.healthHandler)

	r.GET("/basket/:id", s.GetBasketHandler)
	r.POST("/basket", s.StoreBasketHandler)
	r.DELETE("/basket/:id", s.DeleteBasketHandler)
	

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) GetBasketHandler(c *gin.Context) {
	id := c.Param("id")

	basket, err := s.repository.GetBasket(id, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func (s *Server) StoreBasketHandler(c *gin.Context) {
	var basket types.ShoppingCart
	if err := c.BindJSON(&basket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.repository.StoreBasket(&basket, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (s *Server) DeleteBasketHandler(c *gin.Context) {
	id := c.Param("id")

	err := s.repository.DeleteBasket(id, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}


//TODO: Implement the DISCOUNT GRPC call