package repository

import (
	"context"
	"hyphen-backend-SISS/entity"
)

type ImageRepository interface {
	Insert(ctx context.Context, image entity.Image) entity.Image

	FindByID(ctx context.Context, id string) (entity.Image, error)
	// FindByAll(ctx context.Context) []entity.Image

	Update(ctx context.Context, image entity.Image) entity.Image
	// UpdateByUUID(ctx context.Context, uuid uuid.UUID) (entity.Image, error)

	Delete(ctx context.Context, image entity.Image)
}
