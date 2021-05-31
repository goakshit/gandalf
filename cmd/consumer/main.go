package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goakshit/gandalf/config"
	"github.com/goakshit/gandalf/internal/persistence"
	"github.com/goakshit/gandalf/internal/service/billing"
	"github.com/goakshit/gandalf/internal/types"
)

func main() {
	conf := config.New()

	a, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": conf.MessageService.Server,
	})
	if err != nil {
		panic(err)
	}

	_, err = a.CreateTopics(
		context.Background(),
		[]kafka.TopicSpecification{{
			Topic:             conf.MessageService.Topic,
			NumPartitions:     1,
			ReplicationFactor: 1}},
	)
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        conf.MessageService.Server,
		"group.id":                 "myGroup",
		"auto.offset.reset":        "earliest",
		"allow.auto.create.topics": true,
	})
	if err != nil {
		panic(err)
	}

	err = c.Subscribe(conf.MessageService.Topic, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	gc := persistence.GetGormClient()

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		var vh types.VehicleDetails
		if err = json.Unmarshal(msg.Value, &vh); err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		vh.ArrivedAt = msg.Timestamp
		billingSvc := billing.NewBillingService(gc)
		if err = billingSvc.CreateVehicleParkingRecord(context.Background(), vh); err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}
		fmt.Println("Successfully created record in database with reg_no: " + vh.RegNo)
	}
}
