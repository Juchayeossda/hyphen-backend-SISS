package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type ImageModel struct {
	ID    uuid.UUID             `json:"id"`
	Image *multipart.FileHeader `form:"image" validate:"multipart-file-image"`
}

type ImageCreateUpdateReturnModel struct {
	ID uuid.UUID `json:"id"`
}

type ImageReadModel struct {
	ImageData []byte
}
