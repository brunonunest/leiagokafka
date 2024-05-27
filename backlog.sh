# Inside the Kafka container shell
kafka-topics --create --topic high-throughput-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
