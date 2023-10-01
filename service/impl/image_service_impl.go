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

	imageEntity := entity.Image{
		ID:        uuid.New(),
		ImageData: extractImage(imageModel),
	}

	// insert
	service.ImageRepository.Insert(ctx, imageEntity)

	// sync model value or entity value
	imageModel.ID = imageEntity.ID
	return imageModel
}

func (service *ImageServiceImpl) FindByID(ctx context.Context, id string) model.ImageReadModel {
	entityImage, err := service.ImageRepository.FindByID(ctx, id)
	exception.PanicLogging(err)
	return model.ImageReadModel{ImageData: entityImage.ImageData}

}

func (service *ImageServiceImpl) Update(ctx context.Context, imageModel model.ImageModel) model.ImageModel {
	common.Validate(imageModel)

	imageEntity := entity.Image{
		ID:        imageModel.ID,
		ImageData: extractImage(imageModel),
	}

	// update
	service.ImageRepository.Update(ctx, imageEntity)
	return imageModel

}

func (service *ImageServiceImpl) Delete(ctx context.Context, id uuid.UUID) {
	image, err := service.ImageRepository.FindByID(ctx, id.String())
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	service.ImageRepository.Delete(ctx, image)
}

func extractImage(imageModel model.ImageModel) []byte {

	// Cast multipart data type to []byte (model -> entity)
	fileContent, err := imageModel.Image.Open()
	exception.PanicLogging(err)
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	exception.PanicLogging(err)

	return fileBytes
}
