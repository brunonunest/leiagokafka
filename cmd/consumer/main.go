package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Message represents the structure of a message consumed from Kafka.
type Message struct {
	Value string
}

func main() {
	// Channel for receiving messages from the consumer goroutine.
	messageChannel := make(chan Message, 100)

	// Number of workers
	numWorkers := 5

	// Start the consumer worker pool.
	for i := 0; i < numWorkers; i++ {
		go consumerWorker(messageChannel)
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
}

// consumerWorker processes messages from the channel.
func consumerWorker(messageChannel chan Message) {
	for message := range messageChannel {
		// Process the consumed message.
		fmt.Println("Consumed message:", message.Value)
	}
}
