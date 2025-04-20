
CREATE TABLE IF NOT EXISTS users (
    id uuid not null PRIMARY KEY,
    first_name text null,
    last_name text null,
    username text null,
    age int null,
    height int null,
    weight int null,
    gender text null,
    exercise_level text null,
    aim text null,
    status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS accounts(
    id uuid not null PRIMARY KEY,
    user_id uuid not null references users(id),
    email text not null,
    password text not null, 
    status text not null,  
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);