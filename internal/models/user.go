package models

import (
	"context"
	"time"

	"github.com/google/uuid"

	"10.0.0.50/tuan.quang.tran/aioz-ads/pkg/v1/db"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error

	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)

	UpdateUser(ctx context.Context, user *User) error
}

type User struct {
	Id           uuid.UUID  `json:"id"`
	Username     string     `json:"username"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	DisplayEmail string     `json:"display_email"`
	Status       Status     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    uuid.UUID  `json:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at"`
	UpdatedBy    uuid.UUID  `json:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at"`
	DeletedBy    *uuid.UUID `json:"deleted_by"`
}

func NewUser(username, firstName, lastName, email string, createdBy uuid.UUID) *User {
	return &User{
		Id:        uuid.New(),
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Status:    ActiveStatus,
		CreatedAt: time.Now(),
		CreatedBy: createdBy,
		UpdatedAt: time.Now(),
		UpdatedBy: createdBy,
		DeletedAt: nil,
		DeletedBy: nil,
	}
}

func ToUser(u db.User) *User {
	return &User{
		Id:           u.ID,
		DisplayEmail: u.DisplayEmail,
		Email:        u.Email,
		Username:     u.Username,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Status:       Status(u.Status),
		CreatedAt:    u.CreatedAt,
		CreatedBy:    u.CreatedBy,
		UpdatedAt:    u.UpdatedAt,
		UpdatedBy:    u.UpdatedBy,
		DeletedAt:    u.DeletedAt,
		DeletedBy:    u.DeletedBy,
	}
}
