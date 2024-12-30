package routes

import (
	"github.com/NetParamad/Go-Dev-101/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	
	api.Get("/", controllers.Index) 					// http://localhost:3000/api
	api.Get("/books", controllers.GetBook)				// http://localhost:3000/api/books
	api.Get("/books/:id", controllers.GetBookById)		// http://localhost:3000/api/books/1
	api.Post("/books", controllers.CreateBook)			// http://localhost:3000/api/books
	api.Put("/books/:id", controllers.UpdateBook) 		// http://localhost:3000/api/books/1
	api.Delete("/books/:id", controllers.DeleteBook) 	// http://localhost:3000/api/books/1
}