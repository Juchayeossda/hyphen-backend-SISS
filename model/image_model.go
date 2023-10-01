package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type ImageModel struct {
	ID    uuid.UUID
	Image *multipart.FileHeader `form:"image" validate:"multipart-file-image"`
}
