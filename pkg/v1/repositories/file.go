package repositories

import (
	"context"
	"time"

	db "10.0.0.50/tuan.quang.tran/aioz-ads/db/generated"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/metrics"
	"10.0.0.50/tuan.quang.tran/aioz-ads/models"
	"github.com/google/uuid"
)

type FileRepository struct {
	db *db.Queries
}

func NewFileRepository(db *db.Queries) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) CreateFile(
	ctx context.Context,
	file *models.File,
) (*models.File, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("CreateFile").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.CreateFile(ctx, db.CreateFileParams{
		ID:         file.ID,
		BelongToID: file.BelongToID,
		FilePath:   file.FilePath,
		FileType:   file.FileType,
		Status:     file.Status,
		CreatedAt:  file.CreatedAt,
		CreatedBy:  file.CreatedBy,
		UpdatedAt:  file.UpdatedAt,
		UpdatedBy:  file.UpdatedBy,
	})
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (r *FileRepository) GetFileByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.File, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetFileByID").
			Observe(time.Since(t).Seconds())
	}()

	file, err := r.db.GetFileById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.ToFile(file), nil
}

func (r *FileRepository) GetFilesByBelongToID(
	ctx context.Context,
	belongToID uuid.UUID,
) ([]*models.File, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("GetFilesByBelongToID").
			Observe(time.Since(t).Seconds())
	}()

	files, err := r.db.GetFilesByBelongToId(ctx, belongToID)
	if err != nil {
		return nil, err
	}

	fileModels := make([]*models.File, len(files))
	for i, file := range files {
		fileModels[i] = models.ToFile(file)
	}

	return fileModels, nil
}

func (r *FileRepository) UpdateFile(
	ctx context.Context,
	file *models.File,
) (*models.File, error) {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("UpdatedAtFileByID").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.UpdateFile(ctx, db.UpdateFileParams{
		ID:         file.ID,
		BelongToID: file.BelongToID,
		FilePath:   file.FilePath,
		FileType:   file.FileType,
		Status:     file.Status,
		UpdatedAt:  file.UpdatedAt,
		UpdatedBy:  file.UpdatedBy,
	})
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (r *FileRepository) DeleteFile(
	ctx context.Context,
	file *models.File,
) error {
	t := time.Now().UTC()
	defer func() {
		metrics.DbMetricsIns.DbSum.WithLabelValues("DeleteFile").
			Observe(time.Since(t).Seconds())
	}()

	err := r.db.DeleteFile(ctx, db.DeleteFileParams{
		ID:        file.ID,
		DeletedAt: file.DeletedAt,
		DeletedBy: file.DeletedBy,
	})
	if err != nil {
		return err
	}

	return nil
}
