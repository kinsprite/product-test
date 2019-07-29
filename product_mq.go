package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// UserCreatingHandler  calling on user creating
type UserCreatingHandler interface {
	OnCreating(userInfo *UserInfo) error
}

var kafkaBrokerURL = "kafka:9092"
var userCreatingHandler UserCreatingHandler

func setUserCreatingHandler(handler UserCreatingHandler) {
	userCreatingHandler = handler
}

func initMQ() {
	url := os.Getenv("KAFKA_BROKER_URL")

	if url != "" {
		kafkaBrokerURL = url
	}

	go readMQLoop()
}

func readMQLoop() {
	// make a new reader that consumes from 'user-creating-topic'
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{kafkaBrokerURL},
		GroupID:        "consumer-group-product",
		Topic:          "topic-user-creating",
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		MaxWait:        time.Second,
		CommitInterval: time.Second, // flushes commits to Kafka every second
	})

	if r == nil {
		log.Println("ERROR    connecting to MQS 'topic-user-creating': FAILED")
	} else {
		log.Println("INFO    connecting to MQS 'topic-user-creating': OK")
	}

	ctx := context.Background()

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			fmt.Println("INFO   MQ reading loop exit")
			break
		}

		fmt.Printf("INFO   message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		if userCreatingHandler != nil {
			var userInfo UserInfo
			json.Unmarshal(m.Value, &userInfo)
			userCreatingHandler.OnCreating(&userInfo)
		}

		r.CommitMessages(ctx, m)
	}

	r.Close()
}
