package main

import (
	"errors"
	"fmt"
	"go-store/pkg/logging"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"go-store/basket/internal/server"
	"go-store/pkg/resilience"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Run() error {
	var ElasticAddress = os.Getenv("ELASTIC_SEARCH")
	hook, err := logging.NewElasticHook([]string{ElasticAddress})
	if err != nil {
		return fmt.Errorf("error creating hook: %w", err)
	}

	logger, err := logging.NewLogger("logs/log.json")
	if err != nil {
		return fmt.Errorf("error creating logger: %w", err)
	}

	logger.AddHook(hook)
	defer hook.Close()

	newServer := server.NewServer(logger)

	go resilience.GracefulShutdown(newServer)

	err = newServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http newServer error: %w", err)
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}
