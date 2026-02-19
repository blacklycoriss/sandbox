package messaging

import (
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	mu      sync.Mutex
	closed  bool
}

func New(url string) (*RabbitMQ, error) {
	r := &RabbitMQ{}
	if err := r.Connect(url); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RabbitMQ) Connect(url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel: %v", err)
		conn.Close()
		return err
	}

	_, err = ch.QueueDeclare("stock-updates", true, false, false, false, nil)
	if err != nil {
		ch.Close()
		conn.Close()
		return err
	}

	r.conn = conn
	r.channel = ch
	r.closed = false

	// Запускаем мониторинг закрытия
	go r.watchClose(url)

	return nil
}

func (r *RabbitMQ) watchClose(url string) {
	for {
		connClose := r.conn.NotifyClose(make(chan *amqp.Error, 1))
		chClose := r.channel.NotifyClose(make(chan *amqp.Error, 1))

		var reason *amqp.Error
		select {
		case reason = <-connClose:
		case reason = <-chClose:
		}

		log.Printf("RabbitMQ connection/channel closed: %v", reason)

		r.mu.Lock()
		r.closed = true
		r.mu.Unlock()

		// backoff + reconnect
		time.Sleep(5 * time.Second) // или экспоненциальный backoff

		log.Println("Reconnecting to RabbitMQ...")
		if err := r.Connect(url); err != nil {
			log.Printf("Reconnect failed: %v", err)
			continue
		}
		log.Println("Reconnected successfully")
	}
}
