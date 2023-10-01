package impl

import (
	"context"
	"errors"
	"hyphen-backend-SISS/entity"
	"hyphen-backend-SISS/repository"
	"hyphen-backend-SISS/system/exception"

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

func (repository *imageRepositoryImpl) FindByID(ctx context.Context, id string) (entity.Image, error) {
	var image entity.Image
	err := repository.DB.WithContext(ctx).Unscoped().Where("image_id = ?", id).First(&image).Error
	if err != nil {
		return entity.Image{}, errors.New("image not found")
	}
	return image, nil
}

func (repository *imageRepositoryImpl) Update(ctx context.Context, image entity.Image) entity.Image {
	err := repository.DB.WithContext(ctx).Where("image_id = ?", image.ID).Updates(&image).Error
	exception.PanicLogging(err)
	return image
}

func (repository *imageRepositoryImpl) Delete(ctx context.Context, image entity.Image) {
	err := repository.DB.WithContext(ctx).Where("image_id = ?", image.ID).Delete(&image).Error
	exception.PanicLogging(err)

}
