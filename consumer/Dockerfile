FROM golang:latest

MAINTAINER Dmytro Grendach (grendach@gmail.com)

# install dependencies
RUN go get github.com/Shopify/sarama
RUN go get github.com/wvanbergen/kafka/consumergroup


#copy app
ADD src/kcons.go /go/src

# build
RUN go build -o /go/bin/kcons /go/src/kcons.go
RUN chmod +x /go/bin/kcons

ENTRYPOINT [ "/go/bin/kcons" ]
