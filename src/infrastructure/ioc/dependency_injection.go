package ioc

import (
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/database"
	"github.com/jailtonjunior94/financialcontrol-events/src/infrastructure/repositories"
)

var (
	SqlConnection     database.ISqlConnection
	AccountRepository repositories.IAccountRepository
)

func New(sqlConnection database.ISqlConnection) {
	/* Database */
	SqlConnection = sqlConnection

	/* Repositories */
	AccountRepository = repositories.NewAccountRepository(SqlConnection)

	/* Services */

	/* Handlers */
}
