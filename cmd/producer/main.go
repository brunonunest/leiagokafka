package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka broker address
	brokerAddress := "kafka:9092"
	// Topic to produce messages to
	topic := "high-throughput-topic"

	// Create a new Kafka writer with the broker address and topic
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	// Create a message to send
	message := kafka.Message{
		Key:   []byte("Key-A"),
		Value: []byte("Help me Obi-Wan Kenobi, you are my only hope! (Imagine if Leia could use Kafka rs)!"),
	}

	// Send the message
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatalf("failed to write messages: %v", err)
	}

	// Log success and close writer
	log.Println("Message sent successfully")
	if err := writer.Close(); err != nil {
		log.Fatalf("failed to close writer: %v", err)
	}
}
