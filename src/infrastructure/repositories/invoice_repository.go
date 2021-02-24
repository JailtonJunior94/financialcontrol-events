package repositories

import (
	"database/sql"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/domain/entities"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/database"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/queries"
)

type IInvoiceRepository interface {
	InvoiceByDate(ch chan<- []entities.Invoice, startDate time.Time, endDate time.Time) (invoices []entities.Invoice)
}

type InvoiceRepository struct {
	Db database.ISqlConnection
}

func NewInvoiceRepository(db database.ISqlConnection) IInvoiceRepository {
	return &InvoiceRepository{db}
}

func (a *InvoiceRepository) InvoiceByDate(ch chan<- []entities.Invoice, startDate time.Time, endDate time.Time) (invoices []entities.Invoice) {
	database := a.Db.Connect()

	if err := database.Select(&invoices, queries.Invoices, sql.Named("startDate", startDate), sql.Named("endDate", endDate)); err != nil {
		ch <- make([]entities.Invoice, 0)
		return nil
	}
	ch <- invoices
	return invoices
}
