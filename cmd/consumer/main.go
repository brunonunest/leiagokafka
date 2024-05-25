package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	// Configuration for Kafka consumer
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "fc2-gokafka-kafka-1:9092", // Kafka broker address
		"client.id":         "goapp-consumer",           // Consumer ID
		"group.id":          "goapp-group2",             // Consumer group ID
		"auto.offset.reset": "earliest",                 // Read from the beginning if no offset is present
	}

	// Create a new Kafka consumer
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer c.Close()

	// Topics to subscribe to
	topics := []string{"test-topic"}
	if err := c.SubscribeTopics(topics, nil); err != nil {
		log.Fatalf("Error subscribing to topics: %v", err)
	}

	// Continuously read messages
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// Errors are typically informational, the consumer will automatically recover
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
