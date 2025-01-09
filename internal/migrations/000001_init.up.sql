CREATE TABLE users (
    id uuid not null PRIMARY KEY,
    email text not null unique,
    display_email text not null,
    first_name text not null,
    last_name text not null,
    username text not null unique,
    status text not null CHECK (status IN ('active', 'inactive', 'deleted')) DEFAULT 'active',
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid not null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid not null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);

CREATE UNIQUE INDEX idx_email ON users (email);
CREATE UNIQUE INDEX idx_username ON users (username);
