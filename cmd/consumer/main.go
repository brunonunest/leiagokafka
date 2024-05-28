package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Message represents the structure of a message consumed from Kafka.
type Message struct {
	Value string
}

func main() {
	// Channel for receiving messages from the consumer goroutine.
	messageChannel := make(chan Message, 100)

	// Context for controlling the lifetime of workers.
	ctx, cancel := context.WithCancel(context.Background())

	// Number of workers
	numWorkers := 5

	// Use a wait group to wait for all consumer workers to finish.
	var wg sync.WaitGroup

	// Start the consumer worker pool.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go consumerWorker(ctx, messageChannel, &wg)
	}

	// Initialize the Kafka consumer.
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group2",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error creating consumer: %v", err)
	}
	defer consumer.Close() // Ensure the consumer is closed when done.

	// Subscribe to the topic.
	topics := []string{"high-throughput-topic"}
	consumer.SubscribeTopics(topics, nil)

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			// Read message from Kafka.
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				message := Message{
					Value: string(msg.Value),
				}
				messageChannel <- message // Send the message to the channel.
			} else {
				log.Printf("error consuming message: %v", err)
			}
		}
	}()

	<-sigChan
	log.Println("Shutting down gracefully...")
	cancel()              // Cancel the context to stop consumer workers
	close(messageChannel) // Close the channel to stop reading
	wg.Wait()             // Wait for all workers to finish
	log.Println("Consumer shutdown complete")
}

// consumerWorker processes messages from the channel.
func consumerWorker(ctx context.Context, messageChannel chan Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case message, ok := <-messageChannel:
			if !ok {
				return // Channel is closed
			}
			// Process the consumed message.
			fmt.Println("Consumed message:", message.Value)
		case <-ctx.Done():
			return // Context is cancelled
		}
	}
}
