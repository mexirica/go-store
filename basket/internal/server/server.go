package server

import (
	"fmt"
	"go-store/basket/internal/database"
	"go-store/basket/internal/database/repository"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type Server struct {
	port       int
	repository repository.Repository
	db         database.Service
	logger     *logrus.Logger
}

func NewServer(logger *logrus.Logger) *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.WithError(err).Fatal("Error parsing PORT")
	}
	db := database.New()
	NewServer := &Server{
		port:       port,
		db:         db,
		repository: repository.NewMongoRepository(db.GetClient(), logger),
		logger:     logger,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
