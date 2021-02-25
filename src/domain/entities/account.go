package entities

import "time"

type Account struct {
	AccountId   string    `db:"AccountId"`
	AccountDate time.Time `db:"AccountDate"`
	Description string    `db:"Description"`
	Value       float64   `db:"Value"`
}

func SumAccounts(accounts []Account) float64 {
	var total float64

	for _, a := range accounts {
		total += a.Value
	}

	return total
}
