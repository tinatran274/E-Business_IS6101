CREATE TABLE IF NOT EXISTS product_categories (
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

CREATE TABLE IF NOT EXISTS products (
    id uuid not null PRIMARY KEY,
    name text not null,
    description text null,
    brand text null,
    origin text null,
    user_guide text null,
    category_id uuid not null references product_categories(id),
    status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);


CREATE TABLE IF NOT EXISTS product_variants (
    id uuid not null PRIMARY KEY,
    product_id uuid not null references products(id),
    description text null,
    color text not null, 
    retail_price float not null,   
    stock int not null,
    status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS files (
    id uuid not null PRIMARY KEY,
    belong_to_id uuid not null,
    file_path text not null,
    file_type text not null,
    status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);