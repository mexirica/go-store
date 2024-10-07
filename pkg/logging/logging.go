package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

type elasticHook struct {
	client *elastic.Client
}

func NewElasticHook(address []string) (*elasticHook, error) {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: address,
	})
	if err != nil {
		log.Fatalf("Erro ao criar cliente Elasticsearch: %v", err)
	}

	return &elasticHook{client: client}, nil
}

func (h *elasticHook) Fire(entry *logrus.Entry) error {
	data := map[string]interface{}{
		"@timestamp": entry.Time.Format(time.RFC3339Nano),
		"level":      entry.Level.String(),
		"msg":        entry.Message,
		"data":       entry.Data,
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erro ao serializar entrada do Logrus: %v", err)
	}

	_, err = h.client.Index("logs", bytes.NewReader(jsonBytes))
	return err
}

func (h *elasticHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
	}
}

func (h *elasticHook) Close() error {
	return nil
}

func LoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		requestURI := c.Request.RequestURI

		logger.WithFields(logrus.Fields{
			"client_ip":   clientIP,
			"method":      method,
			"status_code": statusCode,
			"request_uri": requestURI,
			"latency_ms":  latencyTime.Milliseconds(),
		}).Info("Request processed")
	}
}

func ErrorMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.WithField("panic_value", r).Error("Panicked")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		}()

		c.Next()
	}
}
