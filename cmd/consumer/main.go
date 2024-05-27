package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group2",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error creating consumer", err.Error())
		return
	}
	defer c.Close()

	topics := []string{"high-throughput-topic"}
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Println("error subscribing to topics", err.Error())
		return
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		} else {
			fmt.Println("error consuming message:", err)
		}
	}
}
