package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Message represents the structure of a message to be sent to Kafka.
type Message struct {
	Key   []byte
	Value []byte
}

func main() {
	// Channel for sending messages to the producer goroutine.
	messageChannel := make(chan Message, 100)

	// Number of workers
	numWorkers := 5

	// Start the producer worker pool.
	for i := 0; i < numWorkers; i++ {
		go producerWorker(messageChannel)
	}

	// Simulate producing messages.
	for i := 0; i < 1000; i++ {
		message := Message{
			Key:   []byte("Key-A"),
			Value: []byte("Help me Obi-Wan Kenobi, you are my only hope! " + time.Now().String()),
		}
		messageChannel <- message
		time.Sleep(10 * time.Millisecond)
	}

	// Close the channel after producing messages.
	close(messageChannel)

	// Wait for a while to ensure all messages are processed
	time.Sleep(5 * time.Second)
}

// producerWorker initializes the Kafka writer and sends messages from the channel.
func producerWorker(messageChannel chan Message) {
	// Kafka broker address and topic.
	brokerAddress := "kafka:9092"
	topic := "high-throughput-topic"

	// Create a new Kafka writer.
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close() // Ensure the writer is closed when done.

	for message := range messageChannel {
		// Write the message to Kafka.
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   message.Key,
			Value: message.Value,
		})
		if err != nil {
			log.Printf("failed to write message: %v", err)
		} else {
			log.Println("Message sent:", string(message.Value))
		}
	}
}
