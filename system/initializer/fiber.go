package initializer

import (
	"hyphen-backend-SISS/system/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {

	return fiber.Config{
		AppName: "hyphen-backend-SISS v1.0.0",

		ErrorHandler: exception.ErrorHandler,
	}
}
