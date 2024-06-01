package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/queries"
	"github.com/kevinkimutai/invoice-management-system/internal/domain"
	"github.com/kevinkimutai/invoice-management-system/internal/utils"
)

func (db *DBAdapter) CreateInvoice(invoice utils.InvoiceParams) (domain.Invoice, error) {
	ctx := context.Background()

	fmt.Println(invoice.UserID)

	// Start Tx
	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return domain.Invoice{}, errors.New("failed to start tx")
	}
	qtx := db.queries.WithTx(tx)

	//Save Invoice
	inv, err := qtx.CreateInvoice(ctx, queries.CreateInvoiceParams{
		InvoiceID:      invoice.InvoiceID,
		UserID:         invoice.UserID,
		CompanyID:      invoice.CompanyID,
		Address:        invoice.Address,
		AccountNumber:  invoice.AccountNumber,
		TotalAmount:    invoice.TotalAmount,
		InvoiceDate:    invoice.InvoiceDate,
		InvoiceDueDate: invoice.InvoiceDueDate,
		InvoiceType:    invoice.InvoiceType,
	})

	if err != nil {
		tx.Rollback(ctx)
		return domain.Invoice{}, err
	}

	//Save InvoiceItem
	var invoiceItems []queries.InvoiceItem

	for _, item := range invoice.Items {
		invoiceItem, err := qtx.CreateInvoiceItem(ctx, queries.CreateInvoiceItemParams{
			InvoiceID: invoice.InvoiceID.String,
			Item:      item.Item,
			Amount:    item.Amount,
		})
		if err != nil {
			tx.Rollback(ctx)
			return domain.Invoice{}, err
		}

		invoiceItems = append(invoiceItems, invoiceItem)
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return domain.Invoice{}, err
	}

	return domain.Invoice{
		ID:             inv.ID,
		InvoiceID:      inv.InvoiceID.String,
		UserID:         inv.UserID.Int64,
		CompanyID:      inv.CompanyID.Int64,
		Address:        inv.Address,
		AccountNumber:  inv.AccountNumber,
		TotalAmount:    utils.ConvertNumericToFloat64(inv.TotalAmount),
		InvoiceDate:    inv.InvoiceDate.Time.String(),
		InvoiceDueDate: inv.InvoiceDueDate.Time.String(),
		InvoiceType:    inv.InvoiceType.String,
		Status:         inv.Status.String,
		CreatedAt:      inv.CreatedAt.Time,
		//Items:          invoiceItems,
	}, nil

}

func (db *DBAdapter) GetInvoice(invoiceID string) (domain.Invoice, error) {
	ctx := context.Background()

	//Convert invoiceID
	var pgtext pgtype.Text
	pgtext.Scan(invoiceID)

	invoice, err := db.queries.GetInvoice(ctx, pgtext)
	if err != nil {
		return domain.Invoice{}, err
	}

	return domain.Invoice{
		ID:             invoice.ID,
		InvoiceID:      invoice.InvoiceID.String,
		UserID:         invoice.UserID.Int64,
		CompanyID:      invoice.CompanyID.Int64,
		Address:        invoice.Address,
		AccountNumber:  invoice.AccountNumber,
		TotalAmount:    utils.ConvertNumericToFloat64(invoice.TotalAmount),
		InvoiceDate:    invoice.InvoiceDate.Time.String(),
		InvoiceDueDate: invoice.InvoiceDueDate.Time.String(),
		Status:         invoice.Status.String,
		InvoiceType:    invoice.InvoiceType.String,
		CreatedAt:      invoice.CreatedAt.Time,
	}, nil
}

func (db *DBAdapter) GetInvoiceData(invoiceID string) (queries.GetInvoiceDataByInvoiceIDRow, error) {
	ctx := context.Background()

	//Convert invoiceID
	var pgtext pgtype.Text
	pgtext.Scan(invoiceID)

	invoice, err := db.queries.GetInvoiceDataByInvoiceID(ctx, pgtext)
	if err != nil {
		return invoice, err
	}

	return invoice, nil

}

func (db *DBAdapter) GetInvoiceItemsData(invoiceID string) ([]queries.GetInvoiceItemDataByInvoiceIDRow, error) {
	ctx := context.Background()

	items, err := db.queries.GetInvoiceItemDataByInvoiceID(ctx, invoiceID)
	if err != nil {
		return items, err
	}

	return items, nil
}
