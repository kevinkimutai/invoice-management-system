package app

import (
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/ports"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

type InvoiceRepo struct {
	db  ports.InvoiceRepoPort
	pdf ports.PDFRepoPort
}

func NewInvoiceRepo(db ports.InvoiceRepoPort, pdf ports.PDFRepoPort) *InvoiceRepo {
	return &InvoiceRepo{
		db:  db,
		pdf: pdf,
	}
}

func (r *InvoiceRepo) CreateInvoice(invoice domain.Invoice) (domain.Invoice, error) {
	//Convert TYpes
	invoiceParams := utils.ConvertInvoiceTypes(invoice)
	inv, err := r.db.CreateInvoice(invoiceParams)

	return inv, err
}

func (r *InvoiceRepo) GetInvoiceByID(invoiceID string) (domain.Invoice, error) {
	invoice, err := r.db.GetInvoice(invoiceID)

	return invoice, err
}

func (r *InvoiceRepo) DownloadInvoice(invoiceID string) error {
	invoice, err := r.db.GetInvoiceData(invoiceID)
	if err != nil {
		return err
	}

	invoiceItems, err := r.db.GetInvoiceItemsData(invoiceID)
	if err != nil {
		return err
	}

	err = r.pdf.GenerateInvoicePDF(invoice, invoiceItems)

	return err
}
