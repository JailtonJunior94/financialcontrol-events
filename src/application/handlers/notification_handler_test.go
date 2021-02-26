package handlers

import (
	"fmt"
	"testing"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/domain/entities"

	"github.com/stretchr/testify/mock"
)

type notificationMock struct {
	mock.Mock
}

func (m *notificationMock) AccountsByDate(ch chan<- []entities.Account, startDate time.Time, endDate time.Time) (accounts []entities.Account) {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %s\n", startDate)

	a := make([]entities.Account, 0)
	ch <- a
	return a
}

func (m *notificationMock) InvoiceByDate(ch chan<- []entities.Invoice, startDate time.Time, endDate time.Time) (invoices []entities.Invoice) {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %s\n", startDate)

	i := make([]entities.Invoice, 0)
	ch <- i
	return i
}

func TestChargeCustomer(t *testing.T) {
	notificationMock := new(notificationMock)

	channelAccount := make(chan []entities.Account)
	notificationMock.On("AccountsByDate", channelAccount, time.Now(), time.Now()).Return(true)

	// next we want to define the service we wish to test
	myService := NotificationHandler{
		AccountRepository: notificationMock,
		InvoiceRepository: notificationMock,
	}
	// and call said method
	myService.SendNotification()

	// at the end, we verify that our myService.ChargeCustomer
	// method called our mocked SendChargeNotification method
	notificationMock.AssertExpectations(t)
}
