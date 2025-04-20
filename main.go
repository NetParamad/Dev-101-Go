package main

import (
	"github.com/NetParamad/Go-Dev-101/config"
	"github.com/NetParamad/Go-Dev-101/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// เชื่อมต่อกับฐานข้อมูล
	config.ConnectDB()

	// สร้างแอปพลิเคชัน Fiber
	app := fiber.New()
	routes.SetupRoutes(app)

	// เริ่มต้นแอปพลิเคชัน
	app.Listen(":3030")
}