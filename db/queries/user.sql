-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1 AND status != 'deleted';

-- name: GetAllUser :one
SELECT *
FROM users 
WHERE status != 'deleted';

-- name: GetUserByEmail :one
SELECT *
FROM users
JOIN accounts ON users.id = accounts.user_id
WHERE accounts.email = $1 AND users.status != 'deleted'
  AND accounts.status != 'deleted';

-- name: CreateUser :exec
INSERT INTO users (
  id, 
  first_name,
  last_name,
  username,
  age,
  height,
  weight,
  gender,
  exercise_level,
  aim,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  ) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
);

-- name: UpdateUser :exec
UPDATE users
SET
  first_name = $2,
  last_name = $3,
  username = $4,
  age = $5,
  height = $6,
  weight = $7,
  gender = $8,
  exercise_level = $9,
  aim = $10,
  status = $11,
  updated_at = $12,
  updated_by = $13
WHERE id = $1 AND status != 'deleted';

-- name: DeleteUser :exec
UPDATE users
SET
  status = 'deleted',
  deleted_at = NOW(),
  deleted_by = $2
WHERE id = $1;

-- name: GetAllUsers :many
SELECT *
FROM users
WHERE status != 'deleted';


