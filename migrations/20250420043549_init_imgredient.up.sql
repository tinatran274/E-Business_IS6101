CREATE TABLE IF NOT EXISTS categories (
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

CREATE TABLE IF NOT EXISTS ingredients (
    id uuid not null PRIMARY KEY,
    name text not null,
    description text null,
    removal float not null, 
    kcal float not null,
    protein float not null,
    lipits float not null,
    glucids float not null, 
    canxi float not null,
    phosphor float not null,
    fe float not null,
    vitamin_a float not null,
    vitamin_b1 float not null,
    vitamin_b2 float not null,
    vitamin_c float not null,
    vitamin_pp float not null,
    beta_caroten float not null,
    category_id uuid not null references categories(id),
    status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS dishes (
    id uuid not null PRIMARY KEY,
    name text not null,
    description text null,
    category_id uuid not null references categories(id),
    status text not null,
    created_at timestamptz not null DEFAULT NOW(),
    created_by uuid null,
    updated_at timestamptz not null DEFAULT NOW(),
    updated_by uuid null,
    deleted_at timestamptz DEFAULT NULL,
    deleted_by uuid DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS recipes (
    dish_id uuid not null references dishes(id),
    ingredient_id uuid not null references ingredients(id),
    unit float not null,
    PRIMARY KEY (dish_id, ingredient_id)
);

