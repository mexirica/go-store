package server

import (
	"fmt"
	"go-store/catalog/internal/database"
	"go-store/catalog/internal/database/repository"
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
	logger *logrus.Logger
}

func NewServer(logger *logrus.Logger) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()
	NewServer := &Server{
		port:       port,
		db:         db,
		logger: logger,
		repository: repository.NewMongoRepository(db.GetClient(),logger),
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
