package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zcubbs/crucible/core/rabbitmq"
	"log"
	"time"
)

type RabbitMQ struct {
	*rabbitmq.RabbitMQ
}

const ExchangeName = "harbinger-events"
const QueueName = "harbinger-events.vega.input"

func (r *RabbitMQ) LaunchEventWorker() {
	_, err := r.Channel.QueueDeclare(
		QueueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}

	// Bind queue to exchange
	err = r.Channel.QueueBind(
		QueueName,    // queue name
		"",           // routing key
		ExchangeName, // exchange
		false,
		nil,
	)

	// Subscribing to Queue for getting messages.
	messages, err := r.Channel.Consume(
		QueueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		log.Println(err)
	}

	var forever chan struct{}

	//go func() {
	//	t := time.Tick(5 * time.Second)
	//	for {
	//		select {
	//		case <-t:
	//			processMessage(messages)
	//		case <-context.Background().Done():
	//			return
	//		}
	//	}
	//}()

	go func() {
		processMessage(messages)
	}()

	log.Println("Successfully started event worker routine")
	<-forever
}

func processMessage(messages <-chan amqp.Delivery) {
	for msg := range messages {
		if msg.Body == nil {
			log.Println("Error, no message body!")
			continue
		}

		event := &HarbingerEvent{}
		err := json.Unmarshal(msg.Body, event)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(event)
	}
}

func (r *RabbitMQ) PublishEvent(event *HarbingerEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(event)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to marshal event: %v", err))
	}

	err = r.Channel.PublishWithContext(ctx,
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return errors.New(fmt.Sprintf("failed to publish event: %v", err))
	}

	return nil
}

type VegaEvent struct {
	EventType string      `json:"eventType"`
	EventTime time.Time   `json:"eventTime"`
	EventData interface{} `json:"eventData"`
}

type HarbingerEvent struct {
	EventType string      `json:"eventType"`
	EventTime time.Time   `json:"eventTime"`
	EventData interface{} `json:"eventData"`
}
