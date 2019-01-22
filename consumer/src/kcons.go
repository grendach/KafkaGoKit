package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"log"
	"os"
	"time"
)

var (
	group         = os.Getenv("KAFKA_GROUP")
	topic         = os.Getenv("KAFKA_TOPIC")
	zookeeperHost = os.Getenv("ZOOKEEPER_HOST")
	zookeeperPort = os.Getenv("ZOOKEEPER_PORT")
	zookeeperURI = zookeeperHost + ":" + zookeeperPort
)

func initConsumer() (*consumergroup.ConsumerGroup, error) {
	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup(group, []string{topic}, []string{zookeeperURI}, config)

	if group == "" {
		fmt.Println("!!! <KAFKA_GROUP> variable is not assigned.")
	}

	if topic == "" {
		fmt.Println("!!! <KAFKA_TOPIC> variable is not assigned.")
	}

	if zookeeperHost == "" {
		fmt.Println("!!! <ZOOKEEPER_HOST> variable is not assigned.")
	}

	if zookeeperPort == "" {
		fmt.Println("!!! <ZOOKEEPER_PORT> variable is not assigned.")
	}

	if err != nil {
		return nil, err
	}
	return cg, err

}

func consume(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			// messages comming through the channel
			// only take messages from subscribed topic
			if msg.Topic != topic {
				continue
			}

			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))

			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart

			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error)
			}
		}
	}
}

func main() {
	// setup sarama group to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// init consumer
	cg, err := initConsumer()
	if err != nil {
		fmt.Println("Error consumer group: ", err.Error())
		os.Exit(1)
	}
	defer cg.Close()

	// run consumer
	consume(cg)
}
