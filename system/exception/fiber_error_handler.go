package exception

import (
	"encoding/json"
	"hyphen-backend-SISS/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	_, isValidationError := err.(ValidationError)

	if isValidationError {
		data := err.Error()
		var messages []map[string]any

		err := json.Unmarshal([]byte(data), &messages)
		PanicLogging(err)
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, isNotFoundError := err.(NotFoundError)
	if isNotFoundError {
		return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}
