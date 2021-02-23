package entities

import "time"

type Account struct {
	AccountId   string    `db:"AccountId"`
	AccountDate time.Time `db:"AccountDate"`
	Description string    `db:"Description"`
	Value       float64   `db:"Value"`
}
