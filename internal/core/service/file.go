package service

import (
	"context"
	"fmt"
	"regexp"

	"github.com/attoleap/objekt/internal/core/domain"
	"github.com/attoleap/objekt/internal/core/port"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type FileService struct {
	log        *zerolog.Logger
	bucketRepo port.BucketRepository
	fileRepo   port.FileRepository
}

var validFileNameRegex = regexp.MustCompile(`^[a-zA-Z]([._-]?[a-zA-Z0-9]{1,})*$`)

// interface guard
var _ port.FileService = (*FileService)(nil)

func NewFileService(log *zerolog.Logger, bucketRepo port.BucketRepository, fileRepo port.FileRepository) *FileService {
	return &FileService{
		log:        log,
		fileRepo:   fileRepo,
		bucketRepo: bucketRepo,
	}
}

func (f *FileService) CreateFile(ctx context.Context, file *domain.File) (*domain.File, error) {
	if !validFileNameRegex.MatchString(file.Name) {
		err := fmt.Errorf("invalid file name: %s", file.Name)
		f.log.Err(err).Msg("invalid file name")
		return nil, err
	}

	if file.Size <= 0 {
		f.log.Error().Int64("file_size", file.Size).Msg("invalid file size")
		return nil, fmt.Errorf("invalid file size: %d", file.Size)
	}

	bucketName := file.BucketName
	bucket, err := f.bucketRepo.GetBucketByName(ctx, bucketName)
	if err != nil {
		f.log.Err(err).Str("bucket_name", bucketName).Msg("failed to get bucket")
		return nil, fmt.Errorf("failed to get bucket: %w", err)
	}

	dbFile, _ := f.fileRepo.GetFileByName(ctx, file.Name, bucket.ID)
	if dbFile != nil {
		if dbFile.IsIdentical(file) {
			f.log.Debug().Str("file_name", file.Name).Str("bucket_name", bucketName).Msg("duplicate file creation attempted")
			return dbFile, nil
		}
		f.log.Error().Str("file_name", file.Name).Str("bucket_name", bucketName).Msg("file already exists")
		return nil, fmt.Errorf("file with name %s already exists in bucket %s", file.Name, bucket.Name)
	}

	return f.fileRepo.CreateFile(ctx, file)
}

func (f *FileService) DeleteFile(ctx context.Context, id string) error {
	fileID, err := uuid.Parse(id)
	if err != nil {
		f.log.Err(err).Str("file_id", id).Msg("invalid file ID")
		return fmt.Errorf("invalid file ID: %w", err)
	}

	_, err = f.fileRepo.GetFileByID(ctx, fileID)
	if err != nil {
		f.log.Err(err).Str("file_id", id).Msg("file not found")
		return fmt.Errorf("failed to get file: %w", err)
	}

	return f.fileRepo.DeleteFile(ctx, fileID)
}

func (f *FileService) GetFile(ctx context.Context, id string) (*domain.File, error) {
	fileID, err := uuid.Parse(id)
	if err != nil {
		f.log.Err(err).Str("file_id", id).Msg("invalid file ID")
		return nil, fmt.Errorf("invalid file ID: %w", err)
	}

	return f.fileRepo.GetFileByID(ctx, fileID)
}

func (f *FileService) GetFilesByBucketID(ctx context.Context, bucketID string) ([]domain.File, error) {
	bucketUUID, err := uuid.Parse(bucketID)
	if err != nil {
		f.log.Err(err).Str("bucket_id", bucketID).Msg("invalid bucket ID")
		return nil, fmt.Errorf("invalid bucket ID: %w", err)
	}

	return f.fileRepo.GetFilesByBucketID(ctx, bucketUUID)
}
