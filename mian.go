package main

import (
	"./routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}