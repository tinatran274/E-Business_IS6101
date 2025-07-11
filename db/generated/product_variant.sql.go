// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: product_variant.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createProductVariant = `-- name: CreateProductVariant :exec
INSERT INTO product_variants (
  id, 
  product_id,
  description,
  color,
  retail_price,
  stock,
  status,
  created_at,
  created_by,
  updated_at,
  updated_by
  )
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
`

type CreateProductVariantParams struct {
	ID          uuid.UUID  `json:"id"`
	ProductID   uuid.UUID  `json:"product_id"`
	Description *string    `json:"description"`
	Color       string     `json:"color"`
	RetailPrice float64    `json:"retail_price"`
	Stock       int32      `json:"stock"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
}

func (q *Queries) CreateProductVariant(ctx context.Context, arg CreateProductVariantParams) error {
	_, err := q.db.Exec(ctx, createProductVariant,
		arg.ID,
		arg.ProductID,
		arg.Description,
		arg.Color,
		arg.RetailPrice,
		arg.Stock,
		arg.Status,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	return err
}

const deleteProductVariant = `-- name: DeleteProductVariant :exec
UPDATE product_variants
SET
  status = 'deleted',
  deleted_at = $2,
  deleted_by = $3
WHERE id = $1 AND status != 'deleted'
`

type DeleteProductVariantParams struct {
	ID        uuid.UUID  `json:"id"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by"`
}

func (q *Queries) DeleteProductVariant(ctx context.Context, arg DeleteProductVariantParams) error {
	_, err := q.db.Exec(ctx, deleteProductVariant, arg.ID, arg.DeletedAt, arg.DeletedBy)
	return err
}

const getProductVariantById = `-- name: GetProductVariantById :one
SELECT id, product_id, description, color, retail_price, stock, status, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
FROM product_variants
WHERE id = $1 AND status != 'deleted'
`

func (q *Queries) GetProductVariantById(ctx context.Context, id uuid.UUID) (ProductVariant, error) {
	row := q.db.QueryRow(ctx, getProductVariantById, id)
	var i ProductVariant
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Description,
		&i.Color,
		&i.RetailPrice,
		&i.Stock,
		&i.Status,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const getProductVariantsByProductId = `-- name: GetProductVariantsByProductId :many
SELECT l.id, l.product_id, l.description, l.color, l.retail_price, l.stock, l.status, l.created_at, l.created_by, l.updated_at, l.updated_by, l.deleted_at, l.deleted_by
FROM product_variants l
WHERE l.product_id = $1 AND l.status != 'deleted'
`

func (q *Queries) GetProductVariantsByProductId(ctx context.Context, productID uuid.UUID) ([]ProductVariant, error) {
	rows, err := q.db.Query(ctx, getProductVariantsByProductId, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductVariant
	for rows.Next() {
		var i ProductVariant
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.Description,
			&i.Color,
			&i.RetailPrice,
			&i.Stock,
			&i.Status,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.DeletedAt,
			&i.DeletedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isColorExist = `-- name: IsColorExist :one
SELECT EXISTS (
  SELECT 1
  FROM product_variants
  WHERE color = $1
    AND product_id = $2
    AND status != 'deleted'
) AS exists
`

type IsColorExistParams struct {
	Color     string    `json:"color"`
	ProductID uuid.UUID `json:"product_id"`
}

func (q *Queries) IsColorExist(ctx context.Context, arg IsColorExistParams) (bool, error) {
	row := q.db.QueryRow(ctx, isColorExist, arg.Color, arg.ProductID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const updateProductVariant = `-- name: UpdateProductVariant :exec
UPDATE product_variants
SET
  product_id = $2,
  description = $3,
  color = $4,
  retail_price = $5,
  stock = $6,
  status = $7,
  updated_at = $8,
  updated_by = $9
WHERE id = $1 AND status != 'deleted'
`

type UpdateProductVariantParams struct {
	ID          uuid.UUID  `json:"id"`
	ProductID   uuid.UUID  `json:"product_id"`
	Description *string    `json:"description"`
	Color       string     `json:"color"`
	RetailPrice float64    `json:"retail_price"`
	Stock       int32      `json:"stock"`
	Status      string     `json:"status"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
}

func (q *Queries) UpdateProductVariant(ctx context.Context, arg UpdateProductVariantParams) error {
	_, err := q.db.Exec(ctx, updateProductVariant,
		arg.ID,
		arg.ProductID,
		arg.Description,
		arg.Color,
		arg.RetailPrice,
		arg.Stock,
		arg.Status,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	return err
}
