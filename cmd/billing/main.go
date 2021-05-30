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

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "80", "caller", log.DefaultCaller)

	// Handlers http
	svc := billing.NewBillingService(gormInstance)
	r := billingTpt.NewHttpServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", "80")
	logger.Log("err", http.ListenAndServe(":80", r))
}
