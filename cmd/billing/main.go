package main

import (
	"fmt"

	"github.com/goakshit/gandalf/persistence"
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
		fmt.Println("Connected to database")
	}
}
