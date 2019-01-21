FROM golang:1.10.4

MAINTAINER Dmytro Grendach (grendach@gmail.com)

#install dependencies

RUN go get github.com/Shopify/sarama

#environment
ENV KAFKA_HOST dev.locahost
ENV KAFKA_PORT 9092
ENV KAFKA_TOPIC grendach

ADD . /app
WORKDIR /app

#build producer
RUN go build -o producer src/*.go


#ENTRYPOINT ["/app/docker-enrypoint.sh"]
