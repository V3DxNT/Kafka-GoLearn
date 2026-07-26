package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	topic := "comments"
	worker, err := connectConsumer([]string{"localhost:29092"})
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sarama consumer up and running!...")

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	msgCount := 0
	doneChan := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Printf("Consumer error: %s\n", err)

			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Received", msgCount, string(msg.Topic), string(msg.Value))

			case <-signChan:
				fmt.Println("Interrupt signal received")
				doneChan <- struct{}{}
			}
		}
	}()

	<-doneChan
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		log.Fatal(err)
	}
}

func connectConsumer(brokerUrls []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	conn, err := sarama.NewConsumer(brokerUrls, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
