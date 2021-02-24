package repositories

import (
	"database/sql"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/domain/entities"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/database"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/queries"
)

type IAccountRepository interface {
	AccountsByDate(ch chan<- []entities.Account, startDate time.Time, endDate time.Time) (accounts []entities.Account)
}

type AccountRepository struct {
	Db database.ISqlConnection
}

func NewAccountRepository(db database.ISqlConnection) IAccountRepository {
	return &AccountRepository{db}
}

func (a *AccountRepository) AccountsByDate(ch chan<- []entities.Account, startDate time.Time, endDate time.Time) (accounts []entities.Account) {
	database := a.Db.Connect()
	if err := database.Select(&accounts, queries.Accounts, sql.Named("startDate", startDate), sql.Named("endDate", endDate)); err != nil {
		ch <- make([]entities.Account, 0)
		return nil
	}
	ch <- accounts
	return accounts
}
