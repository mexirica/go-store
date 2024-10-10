package messagebus

type MessageBus interface {
    Publish(event Event) error
    Subscribe() chan Event
}

type Event struct {
    Type string      `json:"type"`
    Data interface{} `json:"data"`
}

type Producer struct {
    bus MessageBus
}

func (p *Producer) ProduceEvent(event Event) error {
    return p.bus.Publish(event)
}

type Consumer struct {
    bus MessageBus
}

func (c *Consumer) GetEvents() <-chan Event {
    return c.bus.Subscribe()
}