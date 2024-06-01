-- name: GetInvoice :one
SELECT * FROM invoices
WHERE invoice_id = $1 LIMIT 1;

-- name: ListInvoices :many
SELECT * FROM invoices
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CreateInvoice :one
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
RETURNING *;

-- name: UpdateInvoice :exec
UPDATE invoices
  set address = $2,
  account_number = $3,
  total_amount=$4,
  invoice_date = $5,
  invoice_due_date = $6,
  status = $7,
  invoice_type = $8

WHERE invoice_id = $1;

-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE invoice_id = $1;


-- name: GetInvoiceDataByInvoiceID :one
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
LIMIT 1;



