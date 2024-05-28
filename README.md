# LeiaGoKafka Project

This project demonstrates a simple high throughput setup system of a Kafka producer and consumer using Go, Docker, and Docker Compose, being able to help Leia find Obi-Wan. It includes configurations for Kafka, Zookeeper and Go using Docker Compose and Go code to produce and consume messages from Kafka topics, using goroutines and channels.

## Project Structure

```plaintext
fc2-gokafka
│
├── cmd
│   ├── consumer
│   │   └── main.go          # Kafka Consumer implementation in Go
│   ├── producer
│   │   └── main.go          # Kafka Producer implementation in Go
├── .gitignore
├── build-librdkafka.sh      # Script to build and install librdkafka
├── docker-compose.yaml      # Docker Compose configuration for Kafka and Zookeeper
├── Dockerfile               # Dockerfile to build the Go application environment
├── go.mod                   # Go module dependencies
```

# Prerequisites:

Docker and Docker Compose installed on your machine.
Go programming language installed.

# Setup:

# Step 1: Clone the Repository

git clone <repository_url>
cd leiagokafka

# Step 2: Build and Start the Docker Containers

docker-compose up --build -d

This command will start Zookeeper, Kafka and Go services in detached mode.

# Step 3: Create the topic in Kafka

Access the Kafka docker container, and create the topic.

docker exec -it leiagokafka-kafka-1 bash

kafka-topics --create --bootstrap-server=localhost:9092 --topic=high-throughput-topic partitions=3

# Step 4: Running the Consumer

Run the consumer from container to receive messages from Kafka created topic.

docker-compose exec goapp ./consumer

# Step 5: Running the Producer

Run the producer from container to send messages to Kafka created topic.

docker-compose exec goapp ./producer
