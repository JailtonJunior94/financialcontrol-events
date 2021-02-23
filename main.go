package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/shared"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	timer := shared.NewTime()

	s := timer.StartDate()
	e := timer.EndDate()
	l := timer.DaysInMonth(s)

	fmt.Printf("s: %s\n", s)
	fmt.Printf("e: %s\n", e)
	fmt.Printf("l: %d\n", l)

	entryID, err := c.AddFunc("0 0 5 * *", handle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entryID: %d\n", entryID)

	c.Run()
}

func handle() {
	fmt.Printf("Executado em: %s\n", time.Now())
}
