#Kafka GO-Kit

Run Kafka and Zookeeper as Docker images on your local or any other environment.

Also you can build and use as a separate container [producer](producer/README.md) and [consumer](consumer/README.md) which are created in Golang


## Set up Kafka on Docker

1. Run Zookeeper:

```sh
$ docker run -d --name zookeeper -p 2181:2181 jplock/zookeeper
```

2. Run Kafka:

```sh
$ docker run -d --name kafka -p 7203:7203 -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=<IP_ADDRESS> -e ZOOKEEPER_IP=<IP_ADDRESS> ches/kafka
```

KAFKA_ADVERTISED_HOST_NAME is the IP address of the machine(my local machine) which Kafka container running. ZOOKEEPER_IP is the Zookeeper container running machines IP.
You can't use you localhost (127.0.0.1) address, in that case zookeeper container won't start.
Use network interface IP address.

3. Create Kafka topic:
```sh
docker run \
--rm ches/kafka kafka-topics.sh \
--create \
--topic grendach \
--replication-factor 1 \
--partitions 1 \
--zookeeper <IP_ADDRESS>:2181
```

4. List Kafka topics:

```sh
docker run \
--rm ches/kafka kafka-topics.sh \
--list \
--zookeeper <IP_ADDRESS>:2181
```

5. Start publisher (creates a producer for `grendach` topic)

```sh
docker run --rm --interactive \
ches/kafka kafka-console-producer.sh \
--topic grendach \
--broker-list <IP_ADDRESS>:9092
```

This producer will take inputs from command line and publish them to Kafka

6. Start reader (creates a consumer for `grendach` topic)

```sh
docker run --rm --interactive \
ches/kafka kafka-console-consumer.sh \
--topic grendach \
--from-beginning \
--bootstrap-server <IP_ADDRESS>:9092
```

This consumer will take the messages from topic and output into the command line.

`from-beginning` parameter specifies to take the messages from the beginning of the topic.

`broker-list <IP_ADDRESS>:9092` specifies the host and port of Zookeeper

## Working with Zookeeper
You can connect ot Zookeeper directly and check topics, brokers, conusmers, etc.

1. Go inside Zookeeper container:
```sh
docker exec -it zookeeper bash
```

Inside the bin directory we can find the commands which are available to manage Zookeeper

2. Connect to Zookeeper server:
```sh
bin/zkCli.sh -server 127.0.0.1:2181
```
Since we are inside the Zookeeper container, we can specify server address as the localhost(127.0.0.1).

3. List root:
```sh
ls /
```

4. List brokers:
```sh
ls /brokers
```

5. List topics:
```sh
ls /brokers/topics
```
6. List consumers:
```sh
ls /consumers
```
7. List consumer owner:
```sh
ls /consumers/console-consumer-1532/owners
```

Above material partially taken from [λ.eranga](https://medium.com/@itseranga) Medium articles.