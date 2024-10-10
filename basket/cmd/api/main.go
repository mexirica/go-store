package main

import (
	"context"
	"errors"
	"fmt"
	"go-store/pkg/logging"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"go-store/basket/internal/server"
)

const ElasticAddress = "http://localhost:9200"

func gracefulShutdown(apiServer *http.Server) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
}

func main() {
	hook, err := logging.NewElasticHook([]string{ElasticAddress})
	if err != nil {
		log.Fatalf("Error creating hook: %v", err)
	}

	logger, err := logging.NewLogger("logs/log.json")
	if err != nil {
		log.Fatalf("Error creating logger: %v", err)
	}

	logger.AddHook(hook)
	defer hook.Close()

	newServer := server.NewServer(logger)

	go gracefulShutdown(newServer)

	err = newServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("http newServer error: %s", err))
	}
}
