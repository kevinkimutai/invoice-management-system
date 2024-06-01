CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    email VARCHAR(50) UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE company (
    company_id BIGSERIAL PRIMARY KEY,
    logo_url VARCHAR NOT NULL,
    user_id BIGINT REFERENCES users(user_id),
    company_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE invoices (
    id BIGSERIAL PRIMARY KEY,
    invoice_id VARCHAR(50) UNIQUE,
    user_id BIGINT REFERENCES users(user_id),
    company_id BIGINT REFERENCES company(company_id),
    address VARCHAR(150) NOT NULL,
    account_number VARCHAR(50) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL CHECK (amount > 0),
    invoice_date DATE NOT NULL,
    invoice_due_date DATE NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('Pending', 'Paid', 'Cancelled')) DEFAULT 'Pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
