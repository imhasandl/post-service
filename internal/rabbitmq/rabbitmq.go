package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// ExchangeName is the default RabbitMQ exchange used for post-related messages.
const (
	ExchangeName = "posts_exchange"
	QueueName    = "notification_service_queue"
	RoutingKey   = "message-service.notification"
)

// RabbitMQ represents a connection to a RabbitMQ message broker.
// It provides methods for publishing and consuming messages.
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbitMQ creates and initializes a new RabbitMQ client.
// It establishes a connection to the RabbitMQ server using the provided connection URL.
func NewRabbitMQ(connectionURL string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		log.Printf("can't connect to rabbit mq: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("can't connect to the channel: %v", err)
		return nil, err
	}

	// Declare the topic exchange
	if err := ch.ExchangeDeclare(
		ExchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	); err != nil {
		log.Printf("failed to declare exchange: %v", err)
		return nil, err
	}

	// Declare the queue for the notification service
	queue, err := ch.QueueDeclare(
		QueueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("failed to declare queue: %v", err)
		return nil, err
	}

	// Bind the queue to the exchange with a wildcard routing key
	if err := ch.QueueBind(
		queue.Name,   // queue name
		"#",          // routing key - match all messages
		ExchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	); err != nil {
		log.Printf("failed to bind queue: %v", err)
		return nil, err
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

// Close closes the RabbitMQ connection and channel.
// It should be called when the RabbitMQ client is no longer needed.
func (r *RabbitMQ) Close() error {
	if r.Channel != nil {
		if err := r.Channel.Close(); err != nil {
			log.Printf("error closing channel: %v", err)
			return err
		}
	}
	if r.Conn != nil {
		if err := r.Conn.Close(); err != nil {
			log.Printf("error closing connection: %v", err)
			return err
		}
	}
	return nil
}
