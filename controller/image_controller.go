package controller

import (
	"hyphen-backend-SISS/model"
	"hyphen-backend-SISS/service"
	"hyphen-backend-SISS/system/exception"

	"github.com/gofiber/fiber/v2"
)

type ImageController struct {
	service.ImageService
}

func NewImageController(imageService *service.ImageService) *ImageController {
	return &ImageController{ImageService: *imageService}
}

func (controller ImageController) Route(app *fiber.App) {
	app.Post("/api/siss/storage/imges/image", controller.Create)
}

func (controller *ImageController) Create(c *fiber.Ctx) error {
	var request model.ImageModel
	var err error
	request.Image, err = c.FormFile("multipart-file-image")
	exception.PanicLogging(err)

	response := controller.ImageService.Create(c.Context(), request)

	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    201,
		Message: "success",
		Data:    response,
	})
}
