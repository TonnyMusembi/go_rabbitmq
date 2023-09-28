package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// RabbitMQ server URL
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"

	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		time.Sleep(1 * time.Minute) // Wait for one minute before retrying

	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		time.Sleep(1 * time.Minute) // Wait for one minute before retrying

	}
	defer ch.Close()

	// Declare a queue
	queueName := "tony_tech"
	_, err = ch.QueueDeclare(
		queueName, // Queue name
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		time.Sleep(1 * time.Minute) // Wait for one minute before retrying

	}

	// Message to publish
	message := "Hello, publishing queue!"

	// Publish the message to the queue
	err = ch.Publish(
		"",        // Exchange
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf("Published message: %s", message)
	time.Sleep(1 * time.Minute)

}
