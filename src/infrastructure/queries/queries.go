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
					[AccountDate] BETWEEN CONVERT(DATETIME, @startDate)
					AND CONVERT(DATETIME, @endDate)`
	Invoices = `SELECT
					CAST([InvoiceId] AS CHAR(36)) InvoiceId,
					[Description],
					[InvoiceControl],
					[InvoiceDate],
					[InvoiceMonth],
					[InvoiceQuantity],
					[InvoiceValue],
					[InvoiceValueTotal],
					CAST([UserId] AS CHAR(36)) UserId,
					CAST([CardId] AS CHAR(36)) CardId,
					CAST([CategoryId] AS CHAR(36)) CategoryId,
					[PurchaseDate]
				FROM
					dbo.Invoices (NOLOCK)
				WHERE
					InvoiceDate BETWEEN CONVERT(DATETIME, @startDate)
				AND CONVERT(DATETIME, @endDate)`
)
