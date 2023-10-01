package repository

import (
	"context"
	"hyphen-backend-SISS/entity"

	"github.com/google/uuid"
)

type ImageRepository interface {
	Insert(ctx context.Context, image entity.Image) entity.Image

	FindByUUID(ctx context.Context, uuid uuid.UUID) (entity.Image, error)
	// FindByAll(ctx context.Context) []entity.Image

	// Update(ctx context.Context, image entity.Image) entity.Image
	UpdateByUUID(ctx context.Context, uuid uuid.UUID) (entity.Image, error)

	DeleteByUUID(ctx context.Context, uuid uuid.UUID) error
}
