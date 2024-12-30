package controllers

import "github.com/gofiber/fiber/v2"

type Books struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price int    `json:"price"`
}


var books []Books


func init() {
	books = []Books{
		{ID: 1, Title: "The Lord of the Rings", Author: "J.R.R Tolkien", Price: 50},
		{ID: 2, Title: "The Hobbit", Author: "J.R.R Tolkien", Price: 40},
		{ID: 3, Title: "The Silmarillion", Author: "J.R.R Tolkien", Price: 30},
		{ID: 4, Title: "The Children of Hurin", Author: "J.R.R Tolkien", Price: 20},
		{ID: 5, Title: "The Fellowship of the Ring", Author: "J.R.R Tolkien", Price: 10},
	}
}

func Index(c *fiber.Ctx) error {
	return c.SendString("Welcome to Go-Dev-101 with Fiber!")
}

func GetBook(c *fiber.Ctx) error {
	return c.JSON(books)
}

func GetBookById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}
	return c.Status(404).SendString("Book not found")
}

func CreateBook(c *fiber.Ctx) error {
	book := new(Books)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	books = append(books, *book)
	return c.JSON(books)
}

func UpdateBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	book := new(Books)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	for i, b := range books {
		if b.ID == id {
			books[i].Title = book.Title
			books[i].Author = book.Author
			books[i].Price = book.Price
			return c.JSON(books)
		}
	}
	return c.Status(404).SendString("Book not found")
}	

func DeleteBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.JSON(books)
		}
	}
	return c.Status(404).SendString("Book not found")
}
