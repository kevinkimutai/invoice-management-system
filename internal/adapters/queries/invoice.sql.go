// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: invoice.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInvoice = `-- name: CreateInvoice :one
INSERT INTO invoices (
  invoice_id,
  user_id,
  company_id,
  address,
  account_number,
  total_amount,
  invoice_date,
  invoice_due_date,
  status,
  invoice_type
) VALUES (
  $1, $2 , $3, $4, $5, $6, $7, $8, COALESCE(NULLIF($9, ''), 'Pending'),$10
)
RETURNING id, invoice_id, user_id, company_id, address, account_number, total_amount, invoice_date, invoice_due_date, status, created_at, invoice_type
`

type CreateInvoiceParams struct {
	InvoiceID      pgtype.Text
	UserID         pgtype.Int8
	CompanyID      pgtype.Int8
	Address        string
	AccountNumber  string
	TotalAmount    pgtype.Numeric
	InvoiceDate    pgtype.Date
	InvoiceDueDate pgtype.Date
	Column9        interface{}
	InvoiceType    pgtype.Text
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRow(ctx, createInvoice,
		arg.InvoiceID,
		arg.UserID,
		arg.CompanyID,
		arg.Address,
		arg.AccountNumber,
		arg.TotalAmount,
		arg.InvoiceDate,
		arg.InvoiceDueDate,
		arg.Column9,
		arg.InvoiceType,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.InvoiceID,
		&i.UserID,
		&i.CompanyID,
		&i.Address,
		&i.AccountNumber,
		&i.TotalAmount,
		&i.InvoiceDate,
		&i.InvoiceDueDate,
		&i.Status,
		&i.CreatedAt,
		&i.InvoiceType,
	)
	return i, err
}

const deleteInvoice = `-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE invoice_id = $1
`

func (q *Queries) DeleteInvoice(ctx context.Context, invoiceID pgtype.Text) error {
	_, err := q.db.Exec(ctx, deleteInvoice, invoiceID)
	return err
}

const getInvoice = `-- name: GetInvoice :one
SELECT id, invoice_id, user_id, company_id, address, account_number, total_amount, invoice_date, invoice_due_date, status, created_at, invoice_type FROM invoices
WHERE invoice_id = $1 LIMIT 1
`

func (q *Queries) GetInvoice(ctx context.Context, invoiceID pgtype.Text) (Invoice, error) {
	row := q.db.QueryRow(ctx, getInvoice, invoiceID)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.InvoiceID,
		&i.UserID,
		&i.CompanyID,
		&i.Address,
		&i.AccountNumber,
		&i.TotalAmount,	
		&i.InvoiceDate,
		&i.InvoiceDueDate,
		&i.Status,
		&i.CreatedAt,
		&i.InvoiceType,
	)
	return i, err
}

const getInvoiceDataByInvoiceID = `-- name: GetInvoiceDataByInvoiceID :one
SELECT
    i.invoice_id,
    c.company_name,
    c.logo_url,
    i.address,
    i.account_number,
    i.total_amount,
    i.invoice_date,
    i.invoice_due_date,
    i.invoice_type
FROM
    invoices i
    JOIN company c ON i.company_id = c.company_id

WHERE
    i.invoice_id = $1
LIMIT 1
`

type GetInvoiceDataByInvoiceIDRow struct {
	InvoiceID      pgtype.Text
	CompanyName    string
	LogoUrl        string
	Address        string
	AccountNumber  string
	TotalAmount    pgtype.Numeric
	InvoiceDate    pgtype.Date
	InvoiceDueDate pgtype.Date
	InvoiceType    pgtype.Text
}

func (q *Queries) GetInvoiceDataByInvoiceID(ctx context.Context, invoiceID pgtype.Text) (GetInvoiceDataByInvoiceIDRow, error) {
	row := q.db.QueryRow(ctx, getInvoiceDataByInvoiceID, invoiceID)
	var i GetInvoiceDataByInvoiceIDRow
	err := row.Scan(
		&i.InvoiceID,
		&i.CompanyName,
		&i.LogoUrl,
		&i.Address,
		&i.AccountNumber,
		&i.TotalAmount,
		&i.InvoiceDate,
		&i.InvoiceDueDate,
		&i.InvoiceType,
	)
	return i, err
}

const listInvoices = `-- name: ListInvoices :many
SELECT id, invoice_id, user_id, company_id, address, account_number, total_amount, invoice_date, invoice_due_date, status, created_at, invoice_type FROM invoices
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListInvoicesParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListInvoices(ctx context.Context, arg ListInvoicesParams) ([]Invoice, error) {
	rows, err := q.db.Query(ctx, listInvoices, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Invoice
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.InvoiceID,
			&i.UserID,
			&i.CompanyID,
			&i.Address,
			&i.AccountNumber,
			&i.TotalAmount,
			&i.InvoiceDate,
			&i.InvoiceDueDate,
			&i.Status,
			&i.CreatedAt,
			&i.InvoiceType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInvoice = `-- name: UpdateInvoice :exec
UPDATE invoices
  set address = $2,
  account_number = $3,
  total_amount=$4,
  invoice_date = $5,
  invoice_due_date = $6,
  status = $7,
  invoice_type = $8

WHERE invoice_id = $1
`

type UpdateInvoiceParams struct {
	InvoiceID      pgtype.Text
	Address        string
	AccountNumber  string
	TotalAmount    pgtype.Numeric
	InvoiceDate    pgtype.Date
	InvoiceDueDate pgtype.Date
	Status         pgtype.Text
	InvoiceType    pgtype.Text
}

func (q *Queries) UpdateInvoice(ctx context.Context, arg UpdateInvoiceParams) error {
	_, err := q.db.Exec(ctx, updateInvoice,
		arg.InvoiceID,
		arg.Address,
		arg.AccountNumber,
		arg.TotalAmount,
		arg.InvoiceDate,
		arg.InvoiceDueDate,
		arg.Status,
		arg.InvoiceType,
	)
	return err
}
