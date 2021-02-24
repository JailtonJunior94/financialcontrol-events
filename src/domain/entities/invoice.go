package entities

import "time"

type Invoice struct {
	InvoiceId         string    `db:"InvoiceId"`
	Description       string    `db:"Description"`
	InvoiceControl    int64     `db:"InvoiceControl"`
	InvoiceDate       time.Time `db:"InvoiceDate"`
	InvoiceMonth      string    `db:"InvoiceMonth"`
	InvoiceQuantity   int       `db:"InvoiceQuantity"`
	InvoiceValue      float64   `db:"InvoiceValue"`
	InvoiceValueTotal float64   `db:"InvoiceValueTotal"`
	UserId            string    `db:"UserId"`
	CardId            string    `db:"CardId"`
	CategoryId        string    `db:"CategoryId"`
	PurchaseDate      time.Time `db:"PurchaseDate"`
}
