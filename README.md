# LeiaGoKafka Project

This project demonstrates a simple setup of a Kafka producer and consumer using Go, Docker, and Docker Compose, being able to help Leia find Obi-Wan. It includes configurations for Kafka and Zookeeper using Docker Compose and Go code to produce and consume messages from Kafka topics.

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

Prerequisites:

Docker and Docker Compose installed on your machine.
Go programming language installed.

# Setup:

Step 1: Clone the Repository

git clone <repository_url>
cd fc2-gokafka

Step 2: Build and Start the Docker Containers

docker-compose up -d
This command will start Zookeeper and Kafka services in detached mode.

Step 3: Running the Producer

Navigate to the producer directory and run the producer to send a message to Kafka.
cd cmd/producer
go run main.go

Step 4: Running the Consumer

Navigate to the consumer directory and run the consumer to read messages from Kafka.
cd cmd/consumer
go run main.go
