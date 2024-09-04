package services

import (
	"BACKEND-GOLANG-MVVM/internal/config"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var receivedMessages = make(chan string)

func StartConsumer(queueName string) {
	// Declare a queue (must match with the producer)
	_, err := config.RabbitMQChannel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := config.RabbitMQChannel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf("Waiting for messages in queue: %s", queueName)
	// Goroutine to process messages
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// BroadcastMessage(string(d.Body)) // Broadcast pesan ke klien
			// receivedMessages <- string(d.Body)

			// Prepare the payload
			// payload := map[string]string{"message": string(d.Body)}
			// payloadBytes, err := json.Marshal(payload)
			var queueMessage QueueMessage
			err := json.Unmarshal(d.Body, &queueMessage)
			if err != nil {
				log.Printf("Failed to marshal payload: %v", err)
				d.Nack(false, true) // Nack the message and requeue
				continue
			}

			// Prepare the payload
			payload := map[string]string{
				"message":  queueMessage.Message,
				"socketId": queueMessage.SocketID, // Include socketId in the payload
			}
			payloadBytes, err := json.Marshal(payload)
			if err != nil {
				log.Printf("Failed to marshal payload: %v", err)
				continue
			}

			// Send the message to Node.js server
			resp, err := http.Post("http://nodejs:3000/trigger", "application/json", bytes.NewBuffer(payloadBytes))
			if err != nil {
				log.Printf("Failed to send message to Node.js: %v", err)
				continue
			}
			resp.Body.Close()

			d.Ack(false) // Send manual acknowledgment
		}
	}()

	log.Printf("Consumer is ready, waiting for messages in queue: %s", queueName)
	select {} // Keep the consumer running
}

// Function to get the last received message from the channel
func GetLastReceivedMessage() string {
	return <-receivedMessages
}
