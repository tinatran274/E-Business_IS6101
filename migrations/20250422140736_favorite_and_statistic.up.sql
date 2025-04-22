CREATE TABLE IF NOT EXISTS favorites (
    dish_id uuid not null references dishes(id),
    user_id uuid not null references users(id),
    value int not null,
    PRIMARY KEY (dish_id, user_id)
);

CREATE TABLE IF NOT EXISTS statistics (
    updated_at date not null,
    user_id uuid not null references users(id),
    morning_calories float not null,
    lunch_calories float not null,
    dinner_calories float not null,
    snack_calories float not null,
    exercise_calories float not null,
    PRIMARY KEY (updated_at, user_id)
);
