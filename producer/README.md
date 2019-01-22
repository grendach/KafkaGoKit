## Kafka producer

Produce (send) message to created <KAFKA_TOPIC>

How to use it:

1. Build Kafka producer docker image:
```sh
$ docker build --build-arg https_proxy="https://10.144.1.10:8080/" -t kprod:0.1.0 .
```

2. Run docker image and pass required environment variables:

* KAFKA_TOPIC
* KAFKA_HOST
* KAFKA_PORT

```sh
$ docker run -it -a stdin -a stdout -e KAFKA_HOST=10.156.54.204 -e KAFKA_PORT=9092 -e KAFKA_TOPIC="greg" kprod:0.1.0
```
If you run docker image for Kafka producer without `stdin` and `stdout` parameters you can't enter messages and your producer will be send empty messages  endlessly!!!
