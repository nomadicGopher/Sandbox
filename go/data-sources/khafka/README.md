## Setup
### Install Kafka and Zookeeper Via Docker Using Bind Mounts

1. **Create a Docker Network**:
   Create a Docker network for Kafka and Zookeeper to communicate:
   ```bash
   docker network create kafka-network
   ```

2. **Run :
   Run Zookeeper using a bind mount to the desired directory:
   ```bash
   docker run -d --name zookeeper \
   --network kafka-network \
   -p 2181:2181 \
   -e ALLOW_ANONYMOUS_LOGIN=yes \
   bitnami/zookeeper:latest
   ```

3. **Create a script to run:**
   * Zookeeper with a Bind Mount & an ARM compatible image (wurstmeister/kafka is not compatible)
   * Kafka with a Bind Mount
   ```bash
   #!/bin/bash
   # Start ZooKeeper and Kafka containers if they are not already running
   
   # Fetch the host IP
   HOST_IP=$(hostname -I | awk '{print $1}')
   
   # Start ZooKeeper
   if [ ! "$(docker ps -q -f name=zookeeper)" ]; then
       if [ "$(docker ps -aq -f status=exited -f name=zookeeper)" ]; then
           # Cleanup if the container exists but is stopped
           docker start zookeeper
       else
           # Run a new ZooKeeper container
           docker run -d --name zookeeper \
              --network kafka-network \
              -p 2181:2181 \
              -e ALLOW_ANONYMOUS_LOGIN=yes \
              bitnami/zookeeper:latest
       fi
   else
       echo "ZooKeeper container is already running."
   fi
   
   # Start Kafka
   if [ ! "$(docker ps -q -f name=kafka)" ]; then
       if [ "$(docker ps -aq -f status=exited -f name=kafka)" ]; then
           # Cleanup if the container exists but is stopped
           docker start kafka
       else
           # Run a new Kafka container
           docker run -d --name kafka \
              --network kafka-network \
              -p 9092:9092 \
              -e KAFKA_ZOOKEEPER_HOSTS=zookeeper:2181 \
              -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 \
              -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://$HOST_IP:9092 \
              -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT \
              -e KAFKA_CFG_PROCESS_ROLES=broker \
              -e KAFKA_CFG_NODE_ID=1 \
              bitnami/kafka:latest
       fi
   else
       echo "Kafka container is already running."
   fi
   ```

**Network and Ports**: Both containers are connected to the `kafka-network`, and the necessary ports are exposed.

## Resources
* [Download](https://kafka.apache.org/downloads)
* [Golang Integration](https://pkg.go.dev/github.com/segmentio/kafka-go)
