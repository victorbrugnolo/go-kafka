## GO KAFKA 🏘🏢
### 📃 Description 
> A simple example of a Kafka Producer and a Kafka Consumer using Golang.

### ⭐ Technologies 
- [Golang 1.16](https://go.dev/)
- [Apache Kafka](https://kafka.apache.org/)
- [librdkafka](https://github.com/edenhill/librdkafka)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### :question: How to run
#### :warning: Prerequisites:

```
docker
docker-compose
```

#### :fast_forward: Running:
```
docker-compose up -d
```

##### Create Kafka topic

```bash
docker exec -it *kafka-container-created-name* bash

kafka-topics --create --topic transactions --bootstrap-server localhost:9092 --partitions 3
```

##### Run consumer

```bash
docker exec -it *go-container-name* bash

go run cmd/consumer/main.go
```

##### Run producer

```bash
docker exec -it *go-container-name* bash

go run cmd/producer/main.go
```
---


