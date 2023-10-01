package service

import (
	"context"
	"hyphen-backend-SISS/model"
)

type ImageService interface {
	Create(ctx context.Context, model model.ImageModel) model.ImageModel
}
