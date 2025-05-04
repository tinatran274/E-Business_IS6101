-- name: CreateFile :exec
INSERT INTO files (
  id, 
  belong_to_id,
  file_path,
  file_type,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  )
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
);

-- name: GetFileById :one
SELECT *
FROM files
WHERE id = $1 AND status != 'deleted';

-- name: GetFilesByBelongToId :many
SELECT l.*
FROM files l
WHERE l.belong_to_id = $1 AND l.status != 'deleted';


-- name: UpdateFile :exec
UPDATE files
SET
  belong_to_id = $2,
  file_path = $3,
  file_type = $4,
  status = $5,
  updated_at = $6,
  updated_by = $7
WHERE id = $1 AND status != 'deleted';

-- name: DeleteFile :exec
UPDATE files
SET
  status = 'deleted',
  deleted_at = $2,
  deleted_by = $3
WHERE id = $1 AND status != 'deleted';