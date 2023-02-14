package main

import (
	"fmt"
	"github.com/zcubbs/crucible/core/rabbitmq"
	"github.com/zcubbs/crucible/harbinger/configs"
	"github.com/zcubbs/crucible/harbinger/internal"
	"log"
)

func init() {
	configs.Bootstrap()
}

func main() {
	fmt.Println("Starting Harbinger...")

	rmq, err := rabbitmq.NewRabbitMQ(&rabbitmq.Config{
		Host:     configs.Config.Harbinger.RabbitMQ.Host,
		Port:     configs.Config.Harbinger.RabbitMQ.Port,
		Username: configs.Config.Harbinger.RabbitMQ.Username,
		Password: configs.Config.Harbinger.RabbitMQ.Password,
		Exchange: internal.ExchangeName,
	})

	if err != nil {
		log.Fatal(fmt.Errorf("NewRabbitMQ %w", err))
	}

	defer rmq.Close()

	r := &internal.RabbitMQ{RabbitMQ: rmq}

	r.LaunchEventWorker()
}
