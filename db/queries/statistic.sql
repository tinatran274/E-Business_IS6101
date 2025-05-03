-- name: CreateStatistic :exec
INSERT INTO statistics (
    updated_at,
    user_id,
    morning_calories,
    lunch_calories,
    dinner_calories,
    snack_calories,
    exercise_calories
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: GetStatisticByUserIdAndDate :one
SELECT * FROM statistics
WHERE user_id = $1 AND updated_at = $2;


-- name: UpdateStatisticByUserIdAndDate :exec
UPDATE statistics
SET
    morning_calories = $3,
    lunch_calories = $4,
    dinner_calories = $5,
    snack_calories = $6,
    exercise_calories = $7
WHERE user_id = $1 AND updated_at = $2;

