package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/environments"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/ioc"
	"github.com/jailtonjunior94/financialcontrol-events/src/shared"

	"github.com/robfig/cron/v3"
)

func main() {
	environments.New()
	ioc.New()

	c := cron.New()

	timer := shared.NewTime()

	s := timer.StartDate()
	e := timer.EndDate()
	l := timer.DaysInMonth(s)

	fmt.Printf("s: %s\n", s)
	fmt.Printf("e: %s\n", e)
	fmt.Printf("l: %d\n", l)

	entryID, err := c.AddFunc(environments.Cron, handle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entryID: %d\n", entryID)

	c.Run()
}

func handle() {
	fmt.Printf("Executado em: %s\n", time.Now())
}
