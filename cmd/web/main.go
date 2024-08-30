package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/thslopes/bff/internal/config"
	"github.com/thslopes/bff/internal/handler"
)

func main() {
	config.LoadQueries()
	config.LoadResources()

	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/:query", handler.QueryHandler)
	app.Post("/:query", handler.QueryHandler)

	log.Fatal(app.Listen(":3000"))
}
