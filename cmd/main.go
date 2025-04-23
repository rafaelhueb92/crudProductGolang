package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/productsGo/internal/handlers"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {

	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Hello the middleware is running!")
		return c.Next()
	})

	app.Get("/healthcheck", healthcheck)

	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProduct)

	log.Fatal(app.Listen(":9000"))
	//fmt.Println("Hello World!")

}
