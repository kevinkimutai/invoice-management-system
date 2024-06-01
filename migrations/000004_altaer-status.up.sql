-- Step 1: Remove the NOT NULL constraint
ALTER TABLE invoices ALTER COLUMN status DROP NOT NULL;

-- Step 2: Ensure the default value remains
ALTER TABLE invoices ALTER COLUMN status SET DEFAULT 'Pending';
