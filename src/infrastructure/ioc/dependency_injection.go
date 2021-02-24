package ioc

import (
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/database"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/repositories"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/telegram"
)

var (
	SqlConnection     database.ISqlConnection
	AccountRepository repositories.IAccountRepository
	InvoiceRepository repositories.IInvoiceRepository
	Telegram          telegram.ITelegram
)

func New(sqlConnection database.ISqlConnection) {
	/* Database */
	SqlConnection = sqlConnection

	/* Repositories */
	AccountRepository = repositories.NewAccountRepository(SqlConnection)
	InvoiceRepository = repositories.NewInvoiceRepository(SqlConnection)

	/* Telegram */
	Telegram = telegram.NewTelegramService()

	/* Services */

	/* Handlers */
}
