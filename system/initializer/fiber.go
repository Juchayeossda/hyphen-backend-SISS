package initializer

import "github.com/gofiber/fiber/v2"

func NewFiberConfiguration() fiber.Config {

	return fiber.Config{
		AppName: "hyphen-backend-SISS v1.0.0",

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return nil
		},
	}
}
