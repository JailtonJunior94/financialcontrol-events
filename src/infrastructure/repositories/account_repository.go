package repositories

import (
	"database/sql"
	"time"

	"github.com/jailtonjunior94/financialcontrol-events/src/domain/entities"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/database"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/queries"
)

type IAccountRepository interface {
	AccountsByDate(startDate time.Time, endDate time.Time) (accounts []entities.Account, err error)
}

type AccountRepository struct {
	Db database.ISqlConnection
}

func NewAccountRepository(db database.ISqlConnection) IAccountRepository {
	return &AccountRepository{db}
}

func (a *AccountRepository) AccountsByDate(startDate time.Time, endDate time.Time) (accounts []entities.Account, err error) {
	database := a.Db.Connect()

	startParam := sql.Named("startDate", startDate.Format("2006-01-02 15:04:05"))
	endParam := sql.Named("endDate", endDate.Format("2006-01-02 15:04:05"))

	if err := database.Select(&accounts, queries.Accounts, startParam, endParam); err != nil {
		return nil, err
	}
	return accounts, nil
}
