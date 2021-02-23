package queries

const (
	Accounts = `SELECT
					CAST([AccountId] AS CHAR(36)) AccountId,
					[AccountDate],
					[Description],
					[Value]
				FROM
					dbo.[Accounts] (NOLOCK)
				WHERE
					[AccountDate] BETWEEN @startDate
				AND @endDate`
	Invoices = `SELECT
					CAST([InvoiceId] AS CHAR(36)) InvoiceId,
					[Description],
					[InvoiceControl],
					[InvoiceDate],
					[InvoiceMonth],
					[InvoiceQuantity],
					[InvoiceValue],
					[InvoiceValueTotal],
					[UserId],
					[CardId],
					[CategoryId],
					[PurchaseDate]
				FROM
					dbo.Invoices (NOLOCK)
				WHERE
					InvoiceDate BETWEEN @startDate
				AND @endDate`
)
