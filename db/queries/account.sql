-- name: GetAccountById :one
SELECT *
FROM accounts
WHERE id = $1 AND status != 'deleted';

-- name: GetAccountByEmail :one
SELECT *
FROM accounts
WHERE email = $1 AND status != 'deleted';

-- name: CreateAccount :exec
INSERT INTO accounts (
  id,
  user_id,
  email,
  password,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
);

-- name: UpdateAccount :exec
UPDATE accounts
SET
  user_id = $2,
  email = $3,
  password = $4,
  status = $5,
  updated_at = $6,
  updated_by = $7
WHERE id = $1 AND status != 'deleted';


-- name: DeleteAccount :exec
UPDATE accounts
SET
  status = 'deleted',
  deleted_at = NOW(),
  deleted_by = $2
WHERE id = $1 AND status != 'deleted';
