package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-store/catalog/internal/types"
	"go-store/pkg/logging"
	"net/http"
)

var Logger *logrus.Logger

type Product = types.Product

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(logging.LoggingMiddleware(Logger))
	r.Use(logging.ErrorMiddleware(Logger))

	r.Group("/admin").GET("/health", s.healthHandler)

	r.POST("/product", s.CreateProductHandler)
	r.GET("/product/:id", s.GetProductHandler)
	r.GET("/products", s.GetProductsHandler)
	r.PATCH("/product/:id", s.UpdateProductHandler)
	r.DELETE("/product/:id", s.DeleteProductHandler)
	r.GET("/products/category/:category", s.GetProductsByCategoryHandler)

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) CreateProductHandler(c *gin.Context) {
	var product Product
	if err := c.BindJSON(&product); err != nil {
		Logger.WithError(err).Error("Error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.repository.CreateProduct(&product, c)
	if err != nil {
		Logger.WithError(err).Error("Error creating product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.WithFields(logrus.Fields{
		"id":          product.Id,
		"name":        product.Name,
		"price":       product.Price,
		"category":    product.Category,
		"description": product.Description,
		"image":       product.ImageFile,
	}).Info("Product created")

	c.JSON(http.StatusCreated, result)
}

func (s *Server) GetProductHandler(c *gin.Context) {
	id := c.Param("id")

	product, err := s.repository.GetProductById(id, c)
	if err != nil {
		Logger.WithError(err).Error("Error getting product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.Info("Product retrieved")
	c.JSON(http.StatusOK, product)
}

func (s *Server) GetProductsHandler(c *gin.Context) {
	products, err := s.repository.GetProducts(c)
	if err != nil {
		Logger.WithError(err).Error("Error getting products")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.Info("Products retrieved")
	c.JSON(http.StatusOK, products)
}

func (s *Server) UpdateProductHandler(c *gin.Context) {
	var params types.UpdateProductParams
	if err := c.BindJSON(&params); err != nil {
		Logger.WithError(err).Error("Error updating product")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := make(map[string]interface{})
	filter["id"] = c.Param("id")

	result, err := s.repository.UpdateProduct(filter, &params, c)
	if err != nil {
		Logger.WithError(err).Error("Error updating product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.WithFields(logrus.Fields{
		"price":       result.Price,
		"category":    result.Category,
		"description": result.Description,
		"image":       result.ImageFile,
	}).Info("Product updated")

	c.JSON(http.StatusOK, result)
}

func (s *Server) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")

	err := s.repository.DeleteProduct(id, c)
	if err != nil {
		Logger.WithError(err).Error("Error deleting product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.Info("Product deleted")
	c.JSON(http.StatusNoContent, nil)
}

func (s *Server) GetProductsByCategoryHandler(c *gin.Context) {
	category := c.Param("category")

	products, err := s.repository.GetProductsByCategory(category, c)
	if err != nil {
		Logger.WithError(err).Error("Error getting products by category")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Logger.Info("Products retrieved by category")
	c.JSON(http.StatusOK, products)
}
