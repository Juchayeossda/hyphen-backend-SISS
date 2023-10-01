package impl

import (
	"context"
	"errors"
	"hyphen-backend-SISS/entity"
	"hyphen-backend-SISS/repository"
	"hyphen-backend-SISS/system/exception"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type imageRepositoryImpl struct {
	*gorm.DB
}

func NewImageRepository(DB *gorm.DB) repository.ImageRepository {
	return &imageRepositoryImpl{DB: DB}
}

func (repository *imageRepositoryImpl) Insert(ctx context.Context, image entity.Image) entity.Image {
	// image.ID = uuid.New() this logic is handled by the service
	err := repository.DB.WithContext(ctx).Create(&image).Error
	exception.PanicLogging(err)
	return image
}

func (repository *imageRepositoryImpl) FindByUUID(ctx context.Context, uuid uuid.UUID) (entity.Image, error) {
	var image entity.Image
	result := repository.DB.WithContext(ctx).Unscoped().Where("image_uuid = ?", uuid).First(&image)
	if result.RowsAffected == 0 {
		return entity.Image{}, errors.New("image not found")
	}
	return image, nil
}

func (repository *imageRepositoryImpl) UpdateByUUID(ctx context.Context, uuid uuid.UUID) (entity.Image, error) {
	var image entity.Image
	err := repository.DB.WithContext(ctx).Where("product_uuid = ?", uuid).Updates(&image).Error
	if err != nil {
		return entity.Image{}, errors.New("image not update")
	}
	return image, nil
}

func (repository *imageRepositoryImpl) DeleteByUUID(ctx context.Context, uuid uuid.UUID) error {
	var image entity.Image
	err := repository.DB.WithContext(ctx).Where("image_uuid = ?", uuid).Delete(&image).Error
	if err != nil {
		return errors.New("image not delete")
	}
	return nil
}
