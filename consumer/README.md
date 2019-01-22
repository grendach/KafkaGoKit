## Kafka consumer

created in `Golang`

1. Build Kafka consumer docker image:
```sh
$ docker build --build-arg https_proxy="https://10.144.1.10:8080/" -t consumer:0.1.0 .
```

2. Run docker image and pass required environment variables:

* KAFKA_GROUP
* KAFKA_TOPIC
* ZOOKEEPER_HOST
* ZOOKEEPER_PORT

```sh
$ docker run -e ZOOKEEPER_HOST=10.156.54.204 -e ZOOKEEPER_PORT=2181 -e KAFKA_TOPIC="greg" -e KAFKA_GROUP="zgroup" consumer:0.1.0
```