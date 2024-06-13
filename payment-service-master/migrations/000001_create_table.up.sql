CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY default gen_random_uuid(),
    reservation_id UUID NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    payment_method TEXT NOT NULL,
    payment_status TEXT NOT NULL
)