package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goakshit/gandalf/config"
)

func main() {
	fmt.Println("Listening on kafka messages")
	conf := config.New()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": conf.MessageService.Server,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	err = c.SubscribeTopics([]string{conf.MessageService.Topic}, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
