package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/domain/entities"
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
	_, err := c.AddFunc(environments.Cron, handle)
	if err != nil {
		log.Fatal(err)
	}
	c.Run()
}

func handle() {
	now := time.Now()
	fmt.Printf("Inicio do Processo em: %s\n", now)

	timer := shared.NewTime()

	startDate := timer.StartDate()
	endDate := timer.EndDate()

	accountRepository := ioc.AccountRepository
	invoiceRepository := ioc.InvoiceRepository

	channelAccount := make(chan []entities.Account)
	channelInvoice := make(chan []entities.Invoice)

	go accountRepository.AccountsByDate(channelAccount, startDate, endDate)
	go invoiceRepository.InvoiceByDate(channelInvoice, startDate, endDate)

	var (
		accounts []entities.Account
		invoices []entities.Invoice
	)

	for i := 0; i < 2; i++ {
		select {
		case chAccount := <-channelAccount:
			accounts = chAccount
		case chInvoice := <-channelInvoice:
			invoices = chInvoice
		}
	}

	for _, a := range accounts {
		fmt.Printf("AccountId: %s\n", a.AccountId)
		fmt.Printf("AccountDate: %s\n", a.AccountDate)
		fmt.Printf("Description: %s\n", a.Description)
		fmt.Printf("Value: %f\n", a.Value)
		fmt.Println("-------")
	}

	for _, i := range invoices {
		fmt.Printf("InvoiceId: %s\n", i.InvoiceId)
		fmt.Printf("Description: %s\n", i.Description)
		fmt.Printf("InvoiceControl: %d\n", i.InvoiceControl)
		fmt.Printf("InvoiceDate: %s\n", i.InvoiceDate)
		fmt.Printf("InvoiceMonth: %s\n", i.InvoiceMonth)
		fmt.Printf("InvoiceQuantity: %d\n", i.InvoiceQuantity)
		fmt.Printf("InvoiceValue: %f\n", i.InvoiceValue)
		fmt.Printf("InvoiceValueTotal: %f\n", i.InvoiceValueTotal)
		fmt.Printf("UserId: %s\n", i.UserId)
		fmt.Printf("CardId: %s\n", i.CardId)
		fmt.Printf("CategoryId: %s\n", i.CategoryId)
		fmt.Printf("PurchaseDate: %s\n", i.PurchaseDate)
		fmt.Println("-------")
	}

	text := `
		O processo levou: %s
	`

	message := fmt.Sprintf(text, time.Since(now))

	if err := ioc.Telegram.SendMessage(message); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Finalização do Processo em: %s\n", time.Since(now))
}
