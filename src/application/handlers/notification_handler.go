package handlers

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/domain/entities"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/repositories"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/telegram"
	"github.com/jailtonjunior94/financialcontrol-events/src/shared"

	"github.com/leekchan/accounting"
)

type INotificationHandler interface {
	SendNotification()
}

type NotificationHandler struct {
	AccountRepository repositories.IAccountRepository
	InvoiceRepository repositories.IInvoiceRepository
	Telegram          telegram.ITelegram
}

func NewNotificationHandler(a repositories.IAccountRepository, i repositories.IInvoiceRepository, t telegram.ITelegram) INotificationHandler {
	return &NotificationHandler{AccountRepository: a, InvoiceRepository: i, Telegram: t}
}

func (n *NotificationHandler) SendNotification() {
	timer := shared.NewTime()
	fmt.Printf("[INFO]: Inicio do Processo em: %s\n", timer.Now)

	accounts, invoices := n.GetAccountsInvoices(timer.StartDate(), timer.EndDate())
	sumInvoices, sumAccounts, percentageAccounts, total := n.Sum(accounts, invoices)
	message := n.Format(sumInvoices, sumAccounts, percentageAccounts, total, timer.Now)

	if err := n.Telegram.SendMessage(message); err != nil {
		log.Println(err)
	}

	fmt.Printf("[INFO]: Finalização do Processo em: %s\n", time.Since(timer.Now))
}

func (n *NotificationHandler) GetAccountsInvoices(startDate, endDate time.Time) (a []entities.Account, i []entities.Invoice) {
	channelAccount := make(chan []entities.Account)
	channelInvoice := make(chan []entities.Invoice)

	go n.AccountRepository.AccountsByDate(channelAccount, startDate, endDate)
	go n.InvoiceRepository.InvoiceByDate(channelInvoice, startDate, endDate)

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

	return accounts, invoices
}

func (n *NotificationHandler) Sum(accounts []entities.Account, invoices []entities.Invoice) (sumInvoices, sumAccounts, percentageAccounts, total float64) {
	sumInvoices = entities.SumInvoices(invoices)
	sumAccounts = entities.SumAccounts(accounts)
	percentageAccounts = math.Round((40.0 / 100.0) * sumAccounts)
	total = (percentageAccounts + sumInvoices)

	return sumInvoices, sumAccounts, percentageAccounts, total
}

func (n *NotificationHandler) Format(sumInvoices, sumAccounts, percentageAccounts, total float64, now time.Time) string {
	format := accounting.Accounting{Symbol: "R$ ", Precision: 2, Thousand: ".", Decimal: ","}

	sumAccountsFormatted := format.FormatMoneyFloat64(sumAccounts)
	sumInvoicesFormatted := format.FormatMoneyFloat64(sumInvoices)
	percentyFormatted := format.FormatMoneyFloat64(percentageAccounts)
	totalFormatted := format.FormatMoneyFloat64(total)
	dateFormatted := fmt.Sprintf("%s de %d", now.Month().String(), now.Year())

	text := `
			 Fechamento do mês: %s

			Despesas da Casa: %s
			40%% das Despesas da Casa: %s
			Cartões: %s
			Total: %s 

			O processo levou %s para executar
	`
	return fmt.Sprintf(text, dateFormatted, sumAccountsFormatted, percentyFormatted, sumInvoicesFormatted, totalFormatted, time.Since(now))
}
