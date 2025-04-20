package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(context.Context, *User) error
	GetUserById(context.Context, uuid.UUID) (*User, error)
	GetAllUser(context.Context) ([]*User, error)
	CreateUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
}

type User struct {
	ID            uuid.UUID  `json:"id"`
	FirstName     *string    `json:"first_name"`
	LastName      *string    `json:"last_name"`
	Username      *string    `json:"username"`
	Age           *int32     `json:"age"`
	Height        *int32     `json:"height"`
	Weight        *int32     `json:"weight"`
	Gender        *string    `json:"gender"`
	ExerciseLevel *string    `json:"exercise_level"`
	Aim           *string    `json:"aim"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	CreatedBy     *uuid.UUID `json:"created_by"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UpdatedBy     *uuid.UUID `json:"updated_by"`
	DeletedAt     *time.Time `json:"deleted_at"`
	DeletedBy     *uuid.UUID `json:"deleted_by"`
}

func NewUser(
	firstName, lastName, username *string,
	age, height, weight *int32,
	gender, exerciseLevel, aim *string,
) *User {
	return &User{
		ID:            uuid.New(),
		FirstName:     firstName,
		LastName:      lastName,
		Username:      username,
		Age:           age,
		Height:        height,
		Weight:        weight,
		Gender:        gender,
		ExerciseLevel: exerciseLevel,
		Aim:           aim,
		Status:        ActiveStatus,
		CreatedAt:     time.Now().UTC(),
		CreatedBy:     nil,
		UpdatedAt:     time.Now().UTC(),
		UpdatedBy:     nil,
		DeletedAt:     nil,
		DeletedBy:     nil,
	}
}

func ToUser(u db.User) *User {
	return &User{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Username:      u.Username,
		Age:           u.Age,
		Height:        u.Height,
		Weight:        u.Weight,
		Gender:        u.Gender,
		ExerciseLevel: u.ExerciseLevel,
		Aim:           u.Aim,
		Status:        u.Status,
		CreatedAt:     u.CreatedAt,
		CreatedBy:     u.CreatedBy,
		UpdatedAt:     u.UpdatedAt,
		UpdatedBy:     u.UpdatedBy,
		DeletedAt:     u.DeletedAt,
		DeletedBy:     u.DeletedBy,
	}
}
