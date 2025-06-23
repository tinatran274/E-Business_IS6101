CREATE TABLE IF NOT EXISTS cart (
    user_id uuid not null references users(id),
    product_variant_id uuid not null references products(id),
    quantity int not null,
    PRIMARY KEY (user_id, product_variant_id)
);

CREATE TABLE IF NOT EXISTS payment_methods (
    id uuid not null PRIMARY KEY,
    name text not null,
    description text null,
    status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS orders (
    id uuid not null primary key default gen_random_uuid(),
    user_id uuid not null references users(id),
    order_date timestamptz not null,
    receiver_name text not null,
    receiver_phone text not null,
    receiver_address text not null,
    shipping_cost float not null,
    payment_method_id uuid not null references payment_methods(id),
    payment_status text not null,
    shipping_status text not null,
    order_status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS order_items (
    order_id uuid not null references orders(id),
    product_variant_id uuid not null references product_variants(id),
    quantity int not null,
    retail_price float not null,
    PRIMARY KEY (order_id, product_variant_id)
);
