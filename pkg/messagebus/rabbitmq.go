package messagebus

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)
var (
	bus MessageBus
)

func Initialize(address,port,fila string) {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", "guest", "guest", address, port)
    conn, err := amqp.Dial(uri)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %s", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %s", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        fila, // nome da fila
        false,    // não durável
        false,    // não excluída automaticamente
        false,    // não exclusiva
        false,    // sem espera
        nil,      // argumentos
    )
    if err != nil {
        log.Fatalf("Failed to declare queue: %s", err)
    }

    bus = NewMessageBus(ch, q.Name)
}

func NewMessageBus(ch *amqp.Channel, queueName string) MessageBus {
    return &messageBus{
        ch:     ch,
        queue:  queueName,
        consumers: make(map[string]chan Event),
    }
}

type messageBus struct {
    ch     *amqp.Channel
    queue  string
    consumers map[string]chan Event
}

func (mb *messageBus) Publish(event Event) error {
    body, _ := json.Marshal(event)
    return mb.ch.Publish("", mb.queue, false, false, amqp.Publishing{
        ContentType: "application/json",
        Body:        body,
    })
}

func (mb *messageBus) Subscribe() chan Event {
    ctx, _ := context.WithCancel(context.Background())
    c := make(chan Event, 100)
    mb.consumers[closureID(ctx)] = c
    go mb.consume(ctx, c)
    return c
}

func (mb *messageBus) consume(ctx context.Context, c chan Event) {
    msgs, err := mb.ch.Consume(mb.queue, "", true, false, false, false, nil)
    if err != nil {
        log.Printf("Failed to register consumer: %v", err)
        return
    }
    defer close(c)

    for d := range msgs {
        event := Event{}
        err := json.Unmarshal(d.Body, &event)
        if err == nil {
            select {
            case c <- event:
            case <-ctx.Done():
                return
            }
        } else {
            log.Printf("Error unmarshaling event: %v", err)
        }
    }
}

func closureID(ctx context.Context) string {
    return fmt.Sprintf("%p", ctx)
}

func NewProducer() *Producer {
	return &Producer{bus: bus}
}

func NewConsumer() *Consumer {
	return &Consumer{bus: bus}
}