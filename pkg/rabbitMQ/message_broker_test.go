package rabbitMQ_test

import (
	"playground/implementation-rabbitMQ-golang/pkg/rabbitMQ"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	// Define the expected arguments for the mocked functions
	expectedQueueName := "test_queue"
	expectedBody := "Test Message"

	// Create the actual instance of the message broker
	broker, err := rabbitMQ.NewRabbitMQBroker("amqp://guest:guest@localhost:5672/", "user", "password")
	assert.Nil(t, err, "Failed to create message broker")

	// Perform the test
	err = broker.Publish(expectedQueueName, expectedBody)
	assert.Nil(t, err, "publish should not return an error")
}

func TestPublishError(t *testing.T) {
	// Define the expected arguments for the mocked function
	expectedQueueName := ""
	expectedBody := ""

	// Create the actual instance of the message broker
	broker, err := rabbitMQ.NewRabbitMQBroker("amqp://guest:guest@localhost:5672/", "user", "password")
	assert.Nil(t, err, "Failed to create message broker")

	// Perform the test
	err = broker.Publish(expectedQueueName, expectedBody)

	// Assert that an error was returned
	assert.NotNil(t, err, "payload cannot be empty")
	assert.NotNil(t, err, "publish should return an error")
}

func TestConsume(t *testing.T) {
	// Set the expectations for the Channel and QueueDeclare functions on the mock channel
	queueName := "test_queue"

	// Create the RabbitMQBroker instance using the mock connection
	broker, err := rabbitMQ.NewRabbitMQBroker("amqp://guest:guest@localhost:5672/", "user", "password")
	assert.Nil(t, err, "Failed to create message broker")

	// Create a flag to check if the handler function is called
	handlerCalled := false
	handler := func(msg string) {
		handlerCalled = true
	}

	// Call the Consume function and verify the results
	err = broker.Consume(queueName, handler)
	assert.NoError(t, err, "Consume should not return an error")
	assert.True(t, handlerCalled, "Handler function should be called")
}
