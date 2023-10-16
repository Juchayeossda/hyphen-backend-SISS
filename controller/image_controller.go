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
	app.Post("/api/siss/images/image", controller.Create)
	app.Get("/api/siss/images/:image_id", controller.FindByID)
	app.Put("/api/siss/images/:image_id", controller.Update)
	app.Delete("/api/siss/images/:image_id", controller.Delete)
}

func (controller *ImageController) Create(c *fiber.Ctx) error {
	var clientRequest model.ImageModel
	var err error

	err = c.BodyParser(&clientRequest)
	exception.PanicLogging(err)

	clientRequest.Image, err = c.FormFile("image")
	exception.PanicLogging(err)

	response := controller.ImageService.Create(c.Context(), clientRequest)

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
	var err error

	request.ID, err = uuid.Parse(c.Params("image_id"))
	exception.PanicLogging(err)

	request.Image, err = c.FormFile("image")
	exception.PanicLogging(err)

	controller.ImageService.Update(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    nil,
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
