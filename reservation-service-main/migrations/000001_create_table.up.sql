CREATE TYPE reservation_status AS ENUM ('busy', 'empty');

CREATE TABLE restaurants (
    id UUID PRIMARY KEY default gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint DEFAULT 0
);

CREATE TABLE reservations (
    id UUID PRIMARY KEY default gen_random_uuid(),
    user_id UUID ,
    restaurant_id UUID references restaurants(id),
    reservation_time TIMESTAMP NOT NULL,
    status reservation_status NOT NULL default 'busy',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint DEFAULT 0
)

CREATE TABLE reservation_orders(
    id UUID PRIMARY KEY default gen_random_uuid(),
    reservation_id UUID references reservations(id),
    menu_item_id UUID references menus(id),
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint DEFAULT 0
)

CREATE TABLE menus (
    id UUID PRIMARY KEY default gen_random_uuid(),
    restaurant_id UUID references restaurants(id),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint DEFAULT 0
)
