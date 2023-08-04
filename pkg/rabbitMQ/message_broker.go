package rabbitMQ

import (
	"errors"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBroker interface {
	Publish(queueName string, body string) error
	Consume(queueName string, handler func(msg string)) error
}

type RabbitMQBroker struct {
	conn *amqp.Connection
}

func NewRabbitMQBroker(url, username, password string) (*RabbitMQBroker, error) {
	conn, err := amqp.DialConfig(url, amqp.Config{
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{
				Username: username,
				Password: password,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &RabbitMQBroker{conn: conn}, nil
}

func (r *RabbitMQBroker) Publish(queueName string, body string) error {
	if err := r.checkPayloadNotEmpty(queueName); err != nil {
		return err
	}

	if err := r.checkPayloadNotEmpty(body); err != nil {
		return err
	}

	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}

func (r *RabbitMQBroker) Consume(queueName string, handler func(msg string)) error {
	if err := r.checkPayloadNotEmpty(queueName); err != nil {
		return err
	}

	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			handler(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit, press CTRL+C")
	<-forever

	return nil
}

func (r *RabbitMQBroker) checkPayloadNotEmpty(payload string) error {
	if payload == "" {
		return errors.New("payload cannot be empty")
	}
	return nil
}
