package main

import (
	"fmt"
	"log"
	"playground/implementation-rabbitMQ-golang/config"
	messagebroker "playground/implementation-rabbitMQ-golang/pkg/rabbitMQ"
)

func main() {
	// Membaca file JSON yang berisi konfigurasi
	config, err := config.ReadConfigFromFile()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	url := fmt.Sprintf("amqp://guest:guest@%s:%s/", config.Endpoint, config.Port)

	broker, err := messagebroker.NewRabbitMQBroker(url, config.AccessKey, config.SecretKey)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ broker: %v", err)
	}

	queueName := "hello"
	err = broker.Publish(queueName, "Hi RabbitMQ!")
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	err = broker.Consume(queueName, func(msg string) {
		log.Printf("Received a message: %s", msg)
	})
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
}
