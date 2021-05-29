package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goakshit/gandalf/config"
	"github.com/goakshit/gandalf/internal/persistence"
)

func main() {

	// Initialise db connection
	gormInstance := persistence.GetGormClient()
	sqlDB, err := gormInstance.DB()
	if err != nil {
		panic("Failed to get postgres db instance: " + err.Error())
	}
	if sqlDB.Ping() != nil {
		panic("Conncetion to database failed")
	} else {
		fmt.Println("Successfully connected to database")
	}

	conf := config.New()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": conf.MessageService.Server,
		"group.id":          "billingGroup",
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
