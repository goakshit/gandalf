package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/goakshit/gandalf/internal/persistence"
	"github.com/goakshit/gandalf/internal/service/billing"
	billingTpt "github.com/goakshit/gandalf/internal/transport/billing"
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

	// conf := config.New()

	// c, err := kafka.NewConsumer(&kafka.ConfigMap{
	// 	"bootstrap.servers": conf.MessageService.Server,
	// 	"group.id":          "billingGroup",
	// 	"auto.offset.reset": "earliest",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// err = c.SubscribeTopics([]string{conf.MessageService.Topic}, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// defer c.Close()

	// go func() {
	// 	for {
	// 		msg, err := c.ReadMessage(-1)
	// 		if err == nil {
	// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	// 		} else {
	// 			// The client will automatically try to recover from all errors.
	// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	// 		}
	// 	}
	// }()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "80", "caller", log.DefaultCaller)

	svc := billing.NewBillingService(gormInstance)
	r := billingTpt.NewHttpServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", "80")
	logger.Log("err", http.ListenAndServe(":80", r))
}
