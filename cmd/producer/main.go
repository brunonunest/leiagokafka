package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	// Context for controlling the lifetime of workers.
	ctx, cancel := context.WithCancel(context.Background())

	// Number of workers
	numWorkers := 5

	// Use a wait group to wait for all producer workers to finish.
	var wg sync.WaitGroup

	// Start the producer worker pool.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go producerWorker(ctx, messageChannel, &wg)
	}

	// Simulate producing messages.
	go func() {
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
	}()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("Shutting down gracefully...")
	cancel()  // Cancel the context to stop producer workers
	wg.Wait() // Wait for all workers to finish
	log.Println("Producer shutdown complete")
}

// producerWorker initializes the Kafka writer and sends messages from the channel.
func producerWorker(ctx context.Context, messageChannel chan Message, wg *sync.WaitGroup) {
	defer wg.Done()

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

	for {
		select {
		case message, ok := <-messageChannel:
			if !ok {
				return // Channel is closed
			}
			// Write the message to Kafka.
			err := writer.WriteMessages(ctx, kafka.Message{
				Key:   message.Key,
				Value: message.Value,
			})
			if err != nil {
				log.Printf("failed to write message: %v", err)
			} else {
				log.Println("Message sent:", string(message.Value))
			}
		case <-ctx.Done():
			return // Context is cancelled
		}
	}
}
