-- name: CreateInvoiceItem :one
INSERT INTO invoice_items (
  invoice_id,
  item,
  amount
) VALUES (
  $1, $2 , $3
)
RETURNING *;

-- name: GetInvoiceItemDataByInvoiceID :many
SELECT item,amount FROM invoice_items WHERE invoice_id = $1;