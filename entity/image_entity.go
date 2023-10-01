package entity

import (
	"github.com/google/uuid"
)

type Image struct {
	ID        uuid.UUID `gorm:"primaryKey;  column:image_id;        type:varchar(36)"`
	ImageData []byte    `gorm:"not null;    column:image_data;  type:BLOB"`
}

func (Image) TableName() string {
	return "tb_image"
}
