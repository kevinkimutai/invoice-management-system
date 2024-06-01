package domain

import (
	"errors"
	"time"
)

type Invoice struct {
	ID             int64         `json:"id"`
	InvoiceID      string        `json:"invoice_id"`
	UserID         int64         `json:"user_id"`
	CompanyID      int64         `json:"company_id"`
	Address        string        `json:"address"`
	AccountNumber  string        `json:"account_number"`
	TotalAmount    float64       `json:"total_amount"`
	InvoiceDate    string        `json:"invoice_date"`
	InvoiceDueDate string        `json:"invoice_due_date"`
	Status         string        `json:"status"`
	CreatedAt      time.Time     `json:"created-at"`
	InvoiceType    string        `json:"invoice_type"`
	Items          []InvoiceItem `json:"items"`
}

type InvoiceItem struct {
	InvoiceItemID int64     `json:"invoice_item_id"`
	InvoiceID     string    `json:"invoice_id"`
	Item          string    `json:"item"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type InvoiceResponse struct {
	StatusCode uint    `json:"status_code"`
	Message    string  `json:"message"`
	Data       Invoice `json:"data"`
}

func NewInvoice(invoice Invoice) (Invoice, error) {
	if invoice.Address == "" || invoice.AccountNumber == "" || invoice.TotalAmount == 0 {
		return invoice, errors.New("missing values in invoice")
	}

	return invoice, nil
}
