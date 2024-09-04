package services

import (
	"BACKEND-GOLANG-MVVM/internal/config"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type QueueMessage struct {
	SocketID string `json:"socketId"`
	Message  string `json:"message"`
}

func PublishMessage(queueName string, socketId string, message string) error {
	// Deklarasikan pesan sebagai struct yang akan dikirim ke queue
	msg := QueueMessage{
		SocketID: socketId,
		Message:  message,
	}

	// Encode pesan sebagai JSON
	msgBody, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to encode message: %v", err)
		return err
	}

	// Declare a queue
	_, err = config.RabbitMQChannel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Failed to declare queue: %v", err)
		return err
	}

	// Publish message to the queue
	err = config.RabbitMQChannel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json", // Update content type to JSON
			Body:        msgBody,
		})
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Published message to queue %s: %s", queueName, string(msgBody))
	return nil
}
