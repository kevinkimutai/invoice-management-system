CREATE TABLE invoice_items(
    invoice_item_id BIGSERIAL PRIMARY KEY,
    invoice_id VARCHAR(50) REFERENCES invoices(invoice_id) ON DELETE CASCADE NOT NULL,
    item varchar(50) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL CHECK (amount > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE invoices ADD invoice_type VARCHAR(50);
ALTER TABLE invoices
RENAME COLUMN amount TO total_amount;
