package models

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"github.com/google/uuid"
)

type FileRepository interface {
	CreateFile(
		ctx context.Context,
		file *File,
	) (*File, error)
	GetFileByID(
		ctx context.Context,
		id uuid.UUID,
	) (*File, error)
	GetFilesByBelongToID(
		ctx context.Context,
		belongToID uuid.UUID,
	) ([]*File, error)
	UpdateFile(
		ctx context.Context,
		file *File,
	) (*File, error)
	DeleteFile(
		ctx context.Context,
		file *File,
	) error
}

type File struct {
	ID         uuid.UUID  `json:"id"`
	BelongToID uuid.UUID  `json:"belong_to_id"`
	FilePath   string     `json:"file_path"`
	FileType   string     `json:"file_type"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  *uuid.UUID `json:"created_by"`
	UpdatedAt  time.Time  `json:"updated_at"`
	UpdatedBy  *uuid.UUID `json:"updated_by"`
	DeletedAt  *time.Time `json:"deleted_at"`
	DeletedBy  *uuid.UUID `json:"deleted_by"`
}

func NewFile(
	id uuid.UUID,
	belongToID uuid.UUID,
	filePath string,
	fileType string,
	status string,
	createdBy *uuid.UUID,
) *File {
	return &File{
		ID:         uuid.New(),
		BelongToID: belongToID,
		FilePath:   filePath,
		FileType:   fileType,
		Status:     status,
		CreatedAt:  time.Now().UTC(),
		CreatedBy:  createdBy,
		UpdatedAt:  time.Now().UTC(),
		UpdatedBy:  createdBy,
	}
}

func ToFile(i db.File) *File {
	return &File{
		ID:         i.ID,
		BelongToID: i.BelongToID,
		FilePath:   i.FilePath,
		FileType:   i.FileType,
		Status:     i.Status,
		CreatedAt:  i.CreatedAt,
		CreatedBy:  i.CreatedBy,
		UpdatedAt:  i.UpdatedAt,
		UpdatedBy:  i.UpdatedBy,
	}
}
