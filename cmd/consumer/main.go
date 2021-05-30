package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goakshit/gandalf/config"
	"github.com/goakshit/gandalf/internal/persistence"
	"github.com/goakshit/gandalf/internal/service/billing"
	"github.com/goakshit/gandalf/internal/types"
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
		fmt.Println("Successfully created record in database")
	}
}
