FROM golang:latest

MAINTAINER Dmytro Grendach (grendach@gmail.com)

# install dependencies
RUN go get github.com/Shopify/sarama

#copy app
ADD src/kprod.go /go/src

# build
RUN go build -o /go/bin/kprod /go/src/kprod.go
RUN chmod +x /go/bin/kprod

ENTRYPOINT [ "/go/bin/kprod" ]
