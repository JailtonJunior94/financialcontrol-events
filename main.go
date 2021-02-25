package main

import (
	"log"

	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/database"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/environments"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/ioc"

	"github.com/robfig/cron/v3"
)

func main() {
	environments.New()

	sqlConnection := database.NewConnection()
	defer sqlConnection.Disconnect()

	ioc.New(sqlConnection)

	c := cron.New()

	_, err := c.AddFunc(environments.Cron, ioc.NotificationHandler.SendNotification)
	if err != nil {
		log.Fatal(err)
	}

	c.Run()
}
