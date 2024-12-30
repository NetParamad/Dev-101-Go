package main

import (
	// "github.com/NetParamad/Go-Dev-101/config"
	"github.com/NetParamad/Go-Dev-101/routes"
	"github.com/gofiber/fiber/v2"
)



func main() {
	// config.DBConfig()
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}