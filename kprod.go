package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
)

const (
	kafkaBroker = "10.129.212.255:31090"
	topic       = "dupa"
)

func main() {
	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	// read comand line output
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		// publish withour goroutine
		publish(msg, producer)

		// publish with goroutine
		// go publish(msg, producer)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	// setups sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config where you can specify parameters of sending messages
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	// prd, err := sarama.NewAsyncProducer([]string{kafkaBroker}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaBroker}, config)

	return prd, err
}

func publish(message string, producer sarama.SyncProducer) {
	// public sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// public async
	// producer.Input() <- $sarama.ProducerMessage{}

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}