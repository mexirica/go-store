package server

import (
	"go-store/basket/internal/types"
	"go-store/pkg/logging"
	"net/http"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "go-store/discount/pkg/grpc"
	"github.com/gin-gonic/gin"
)


func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(logging.LoggingMiddleware(s.logger))
	r.Use(logging.ErrorMiddleware(s.logger))

	r.Group("/admin").GET("/health", s.healthHandler)

	r.GET("/basket/:id", s.GetBasketHandler)
	r.POST("/basket", s.StoreBasketHandler)
	r.DELETE("/basket/:id", s.DeleteBasketHandler)

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	s.logger.Info("Health endpoint called")
	result := s.db.Health()
	s.logger.WithField("health_result", result).Info("Health endpoint response")
	c.JSON(http.StatusOK, result)
}

func (s *Server) GetBasketHandler(c *gin.Context) {
	id := c.Param("id")
	s.logger.WithField("basket_id", id).Info("Request to get shopping basket")

	basket, err := s.repository.GetBasket(id, c)
	if err != nil {
		s.logger.WithError(err).WithField("basket_id", id).Error("Error getting shopping basket")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	s.logger.WithFields(logrus.Fields{
		"basket_id": id,
		"items":     len(basket.Items),
	}).Info("Shopping basket retrieved successfully")
	c.JSON(http.StatusOK, basket)
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	s.logger.Info("Hello World endpoint called")
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func (s *Server) StoreBasketHandler(c *gin.Context) {
	var basket types.ShoppingCart
	if err := c.BindJSON(&basket); err != nil {
		s.logger.WithError(err).Error("Error deserializing shopping basket data")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	basket.TotalPrice = basket.GetTotalPrice()

	err := deductDiscount(&basket, c, s.logger)
	if err != nil {
		s.logger.WithError(err).Error("Error deducting discount")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	s.logger.WithFields(logrus.Fields{
		"user_name":   basket.UserName,
		"items":      len(basket.Items),
		"totalPrice": basket.TotalPrice,
	}).Info("Attempting to store shopping basket")

	result, err := s.repository.StoreBasket(&basket, c)
	if err != nil {
		s.logger.WithError(err).WithField("user_name", basket.UserName).Error("Error storing shopping basket")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	s.logger.WithFields(logrus.Fields{
		"user_name": basket.UserName,
		"result":   result,
	}).Info("Shopping basket stored successfully")
	c.JSON(http.StatusCreated, result)
}

func (s *Server) DeleteBasketHandler(c *gin.Context) {
	id := c.Param("id")
	s.logger.WithField("basket_id", id).Info("Request to delete shopping basket")

	err := s.repository.DeleteBasket(id, c)
	if err != nil {
		s.logger.WithError(err).WithField("basket_id", id).Error("Error deleting shopping basket")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	s.logger.WithField("basket_id", id).Info("Shopping basket deleted successfully")
	c.JSON(http.StatusNoContent, nil)
}

func deductDiscount(cart *types.ShoppingCart, ctx *gin.Context, logger *logrus.Logger) error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	 logger.Fatal("failed to connect to gRPC server at localhost:50051: %v", err)
	 return err
	}
	defer conn.Close()
	c := pb.NewDiscountProtoServiceClient(conn)
   
	r, err := c.GetDiscount(ctx, &pb.GetDiscountRequest{})
	if err != nil {
	 logrus.Fatal("error calling function SayHello: %v", err)
	 return err
	}

	cart.TotalPrice = cart.GetTotalPrice() - float32(r.Amount)
	return nil
}