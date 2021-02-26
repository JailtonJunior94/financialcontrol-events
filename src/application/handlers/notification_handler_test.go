package handlers_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/application/handlers"
	"github.com/jailtonjunior94/financialcontrol-events/src/domain/entities"

	"github.com/stretchr/testify/mock"
)

type NotificationHandlerMock struct {
	mock.Mock
}

func (m *NotificationHandlerMock) SendMessage(message string) error {
	fmt.Println("Mocked SendMessage function")
	fmt.Printf("Params: %s\n", message)

	args := m.Called(message)
	return args.Error(0)
}

func (m *NotificationHandlerMock) AccountsByDate(ch chan<- []entities.Account, startDate time.Time, endDate time.Time) (accounts []entities.Account) {
	fmt.Println("Mocked AccountsByDate function")
	fmt.Printf("Params: %s, %s\n", startDate, endDate)

	a := make([]entities.Account, 0)
	ch <- a

	return a
}

func (m *NotificationHandlerMock) InvoiceByDate(ch chan<- []entities.Invoice, startDate time.Time, endDate time.Time) (invoices []entities.Invoice) {
	fmt.Println("Mocked InvoiceByDate function")
	fmt.Printf("Params: %s, %s\n", startDate, endDate)

	i := make([]entities.Invoice, 0)
	ch <- i

	return i
}

func TestSendNotification(t *testing.T) {
	/* Arrange */
	notificationMock := new(NotificationHandlerMock)

	notification := handlers.NotificationHandler{
		AccountRepository: notificationMock,
		InvoiceRepository: notificationMock,
		Telegram:          notificationMock,
	}

	notificationMock.On("SendMessage", mock.Anything).Return(nil)

	/* Act */
	notification.SendNotification()

	/* Assert */
	notificationMock.AssertExpectations(t)
}

func TestSendNotificationError(t *testing.T) {
	/* Arrange */
	notificationMock := new(NotificationHandlerMock)

	notification := handlers.NotificationHandler{
		AccountRepository: notificationMock,
		InvoiceRepository: notificationMock,
		Telegram:          notificationMock,
	}

	notificationMock.On("SendMessage", mock.Anything).Return(errors.New("Deu ruim"))

	/* Act */
	notification.SendNotification()

	/* Assert */
	notificationMock.AssertExpectations(t)
}
