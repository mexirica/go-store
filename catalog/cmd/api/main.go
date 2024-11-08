package main

import (
	"errors"
	"fmt"
	"go-store/catalog/internal/server"
	"go-store/pkg/logging"
	"go-store/pkg/resilience"
	"log"
	"net/http"
)

const ElasticAddress = "http://localhost:9200"

func main() {
	hook, err := logging.NewElasticHook([]string{ElasticAddress})
	if err != nil {
		log.Fatalf("Error creating hook: %v", err)
	}

	logger, err := logging.NewLogger("logs/api.log")
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	logger.AddHook(hook)
	defer hook.Close()

	newServer := server.NewServer(logger)

	go resilience.GracefulShutdown(newServer)

	err = newServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("http newServer error: %s", err))
	}
}
