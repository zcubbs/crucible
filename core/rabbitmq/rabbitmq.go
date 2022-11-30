package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ ...
type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Exchange string
}

// NewRabbitMQ instantiates the RabbitMQ instances using configuration defined in environment variables.
func NewRabbitMQ(conf *Config) (*RabbitMQ, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("amqp.Dial %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("conn.Channel %w", err)
	}

	err = ch.ExchangeDeclare(
		conf.Exchange, // name
		"topic",       // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("ch.ExchangeDeclare %w", err)
	}

	if err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		return nil, fmt.Errorf("ch.Qos %w", err)
	}

	// XXX: Dead Letter Exchange will be implemented in future episodes

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
	}, nil
}

// Close ...
func (r *RabbitMQ) Close() {
	err := r.Connection.Close()
	if err != nil {
		panic(err)
	}
}
