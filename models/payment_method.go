package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type PaymentMethodRepository interface {
	GetPaymentMethods(
		ctx context.Context,
		filter FilterParams,
	) ([]*PaymentMethod, error)
	CountPaymentMethods(
		ctx context.Context,
	) (int, error)
	GetPaymentMethodByID(ctx context.Context, id uuid.UUID) (*PaymentMethod, error)
}

type PaymentMethod struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   *uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   *uuid.UUID `json:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at"`
	DeletedBy   *uuid.UUID `json:"deleted_by"`
}

func NewPaymentMethod(
	name string,
	description *string,
	createdBy *uuid.UUID,
) *PaymentMethod {
	return &PaymentMethod{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Status:      ActiveStatus,
		CreatedAt:   time.Now().UTC(),
		CreatedBy:   createdBy,
		UpdatedAt:   time.Now().UTC(),
		UpdatedBy:   createdBy,
	}
}

func ToPaymentMethod(d db.PaymentMethod) *PaymentMethod {
	return &PaymentMethod{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt,
		CreatedBy:   d.CreatedBy,
		UpdatedAt:   d.UpdatedAt,
		UpdatedBy:   d.UpdatedBy,
		DeletedAt:   d.DeletedAt,
		DeletedBy:   d.DeletedBy,
	}
}
