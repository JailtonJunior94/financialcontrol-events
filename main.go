package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/database"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/environments"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/ioc"
	"github.com/jailtonjunior94/financialcontrol-events/src/shared"

	"github.com/robfig/cron/v3"
)

func main() {
	environments.New()

	sqlConnection := database.NewConnection()
	defer sqlConnection.Disconnect()

	ioc.New(sqlConnection)

	c := cron.New()

	timer := shared.NewTime()

	startDate := timer.StartDate()
	endDate := timer.EndDate()

	a := ioc.AccountRepository
	b, err := a.AccountsByDate(startDate, endDate)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range b {
		fmt.Printf("AccountId: %s\n", i.AccountId)
		fmt.Printf("AccountDate: %s\n", i.AccountDate)
		fmt.Printf("Description: %s\n", i.Description)
		fmt.Printf("Value: %f\n", i.Value)
		fmt.Println("-------")
	}

	entryID, err := c.AddFunc(environments.Cron, handle)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID da Função: %d\n", entryID)
	fmt.Println("Aguardando próxima execução...")

	c.Run()
}

func handle() {
	fmt.Printf("Executado em: %s\n", time.Now())
}
