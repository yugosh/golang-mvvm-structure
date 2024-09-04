package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

func InitRabbitMQ() {
	var err error

	log.Println("rabbitURL", os.Getenv("RABBITMQ_URL"))

	// Buat koneksi ke RabbitMQ
	RabbitMQConn, err = amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// Buat channel di RabbitMQ
	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	log.Println("Connected to RabbitMQ")
}

func CloseRabbitMQ() {
	RabbitMQChannel.Close()
	RabbitMQConn.Close()
}
