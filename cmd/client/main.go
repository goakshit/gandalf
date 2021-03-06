package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goakshit/gandalf/config"
)

func main() {

	conf := config.New()
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	rand.Seed(time.Now().UnixNano())
	// Produce messages to topic (asynchronously)
	topic := conf.MessageService.Topic
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(fmt.Sprintf(`{"reg_no": "JK02JY%04d", "type": "two"}`, rand.Intn(9999))),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
