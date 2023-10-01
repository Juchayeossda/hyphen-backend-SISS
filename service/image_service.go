package service

import (
	"context"
	"hyphen-backend-SISS/model"

	"github.com/google/uuid"
)

type ImageService interface {
	Create(ctx context.Context, model model.ImageModel) model.ImageModel
	FindByID(ctx context.Context, id string) model.ImageReadModel
	Update(ctx context.Context, imageModel model.ImageModel) model.ImageModel
	Delete(ctx context.Context, id uuid.UUID)
}
