package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type ImageModel struct {
	ID    uuid.UUID             `json:"id"`
	Image *multipart.FileHeader `form:"image" validate:"multipart-file-image" json:"image"`
}

type ImageReadModel struct {
	ImageData []byte
}
