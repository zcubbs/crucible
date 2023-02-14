package internal

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zcubbs/crucible/core/rabbitmq"
	"log"
	"time"
)

type RabbitMQ struct {
	*rabbitmq.RabbitMQ
}

const ExchangeName = "vega-events"
const QueueName = "vega-events.harbinger.input"

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

		event := &VegaEvent{}
		err := json.Unmarshal(msg.Body, event)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(event)
	}
}

type VegaEvent struct {
	EventType string      `json:"eventType"`
	EventTime time.Time   `json:"eventTime"`
	EventData interface{} `json:"eventData"`
}
