CREATE TABLE IF NOT EXISTS sales (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NULL,
    product VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    total NUMERIC(12, 2) GENERATED ALWAYS AS (quantity * price) STORED,
    amount_received NUMERIC(12, 2) NOT NULL CHECK (amount_received >= 0),
    change_amount NUMERIC(12, 2) GENERATED ALWAYS AS (amount_received - (quantity * price)) STORED,
    transaction_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_debt BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_sales_transaction_date ON sales(transaction_date);
CREATE INDEX IF NOT EXISTS idx_sales_created_at ON sales(created_at);