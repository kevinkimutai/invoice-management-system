package ports

import "github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"

type PDFRepoPort interface {
	GenerateInvoicePDF(queries.GetInvoiceDataByInvoiceIDRow, []queries.GetInvoiceItemDataByInvoiceIDRow) error
}
