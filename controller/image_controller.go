package controller

import (
	"hyphen-backend-SISS/model"
	"hyphen-backend-SISS/service"
	"hyphen-backend-SISS/system/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ImageController struct {
	service.ImageService
}

func NewImageController(imageService *service.ImageService) *ImageController {
	return &ImageController{ImageService: *imageService}
}

func (controller ImageController) Route(app *fiber.App) {
	app.Post("/api/siss/storages/images/image", controller.Create)
	app.Get("/api/siss/storages/images/:image_id", controller.FindByID)
	app.Put("/api/siss/storages/images/:image_id", controller.Update)
	app.Delete("/api/siss/storages/images/:image_id", controller.Delete)
}

func (controller *ImageController) Create(c *fiber.Ctx) error {
	var request model.ImageModel
	var err error

	request.Image, err = c.FormFile("multipart-file-image")
	exception.PanicLogging(err)

	response := controller.ImageService.Create(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "Success",
		Data:    response,
	})
}

func (controller *ImageController) FindByID(c *fiber.Ctx) error {
	result := controller.ImageService.FindByID(c.Context(), c.Params("image_id"))
	return c.Status(fiber.StatusOK).Send(result.ImageData)
}

func (controller *ImageController) Update(c *fiber.Ctx) error {
	var request model.ImageModel

	// TODO: change logic
	imageID, err := uuid.Parse(c.Params("image_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    err.Error(),
		})
	}

	request.ID = imageID

	request.Image, err = c.FormFile("multipart-file-image")
	exception.PanicLogging(err)

	response := controller.ImageService.Update(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller *ImageController) Delete(c *fiber.Ctx) error {
	imageID, err := uuid.Parse(c.Params("image_id"))
	exception.PanicLogging(err)

	controller.ImageService.Delete(c.Context(), imageID)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}
