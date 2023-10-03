package main

import (
	"flag"
	"fmt"
	"hyphen-backend-SISS/controller"
	"hyphen-backend-SISS/system"
	"hyphen-backend-SISS/system/initializer"

	repository "hyphen-backend-SISS/repository/impl"
	service "hyphen-backend-SISS/service/impl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var port = `:8080`

func init() {
	portPtr := flag.Int("p", 8080, "Enter the port")
	flag.Parse()
	port = fmt.Sprintf(":%d", *portPtr)
}

func main() {

	config := system.NewConfig()
	database := initializer.NewDatabase(config)

	imageRepository := repository.NewImageRepository(database)

	imageService := service.NewImageService(&imageRepository)

	imageController := controller.NewImageController(&imageService)

	app := fiber.New(initializer.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	imageController.Route(app)

	app.Listen(port)

}
