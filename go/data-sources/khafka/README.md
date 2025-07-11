## Setup
### Install Kafka and Zookeeper Via Docker Using Bind Mounts

1. **Create a Docker Network**:
   Create a Docker network for Kafka and Zookeeper to communicate:
   ```bash
   docker network create kafka-network
   ```

2. **Run Zookeeper with a Bind Mount**:
   Run Zookeeper using a bind mount to the desired directory:
   ```bash
   docker run -d --name zookeeper \
     --network kafka-network \
     -p 2181:2181 \
     -v /ssd/docker/volumes/zookeeper_data:/data \
     wurstmeister/zookeeper
   ```

3. **Create a script in ~/scripts & run Kafka with a Bind Mount**:
   Next, run Kafka with a bind mount:
   ```bash
   docker run -d --name kafka \
     --network kafka-network \
     -p 9092:9092 \
     -v /ssd/docker/volumes/kafka_data:/kafka \
     -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
     -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://<your_host_ip>:9092 \
     -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
     -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 \
     wurstmeister/kafka
   ```

**Network and Ports**: Both containers are connected to the `kafka-network`, and the necessary ports are exposed.

## Resources
* [Download](https://kafka.apache.org/downloads)
* [Golang Integration](https://pkg.go.dev/github.com/segmentio/kafka-go)
