package impl

import (
	"context"
	"hyphen-backend-SISS/entity"
	"hyphen-backend-SISS/model"
	"hyphen-backend-SISS/repository"
	"hyphen-backend-SISS/service"
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

func (service *ImageServiceImpl) Create(ctx context.Context, imageModel model.ImageModel) model.ImageCreateUpdateReturnModel {

	imageEntity := entity.Image{
		ID:        uuid.New(),
		ImageData: extractImage(imageModel),
	}

	// insert
	service.ImageRepository.Insert(ctx, imageEntity)

	return model.ImageCreateUpdateReturnModel{ID: imageEntity.ID}
}

func (service *ImageServiceImpl) FindByID(ctx context.Context, id string) model.ImageReadModel {
	entityImage, err := service.ImageRepository.FindByID(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return model.ImageReadModel{ImageData: entityImage.ImageData}

}

func (service *ImageServiceImpl) Update(ctx context.Context, imageModel model.ImageModel) model.ImageCreateUpdateReturnModel {

	imageEntity := entity.Image{
		ID:        imageModel.ID,
		ImageData: extractImage(imageModel),
	}

	// update
	service.ImageRepository.Update(ctx, imageEntity)

	return model.ImageCreateUpdateReturnModel{ID: imageEntity.ID}

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
