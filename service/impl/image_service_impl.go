package impl

import (
	"context"
	"hyphen-backend-SISS/entity"
	"hyphen-backend-SISS/model"
	"hyphen-backend-SISS/repository"
	"hyphen-backend-SISS/service"
	"hyphen-backend-SISS/system/common"
	"hyphen-backend-SISS/system/exception"
	"io"

	"github.com/google/uuid"
)

type ImageServiceImpl struct {
	repository.ImageRepository
}

func NewImageService(imageRepository *repository.ImageRepository) service.ImageService {
	return &ImageServiceImpl{ImageRepository: *imageRepository}
}

func (service *ImageServiceImpl) Create(ctx context.Context, imageModel model.ImageModel) model.ImageModel {
	common.Validate(imageModel)

	// Cast multipart data type to []byte (model -> entity)
	fileContent, err := imageModel.Image.Open()
	exception.PanicLogging(err)
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	exception.PanicLogging(err)

	imageEntity := entity.Image{
		ID:        uuid.New(),
		ImageData: fileBytes,
	}

	// insert
	service.ImageRepository.Insert(ctx, imageEntity)

	// sync model value or entity value
	imageModel.ID = imageEntity.ID
	return imageModel
}
