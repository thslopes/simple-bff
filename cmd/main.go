package main

import (
	"fmt"
	"log"

	"github.com/thslopes/bff/apicall"
	"github.com/thslopes/bff/setup"

	"github.com/gofiber/fiber/v3"
)

func main() {
	queries, err := setup.LoadQueries()
	if err != nil {
		fmt.Println(err)
		return
	}
	apicall.Queries = queries

	resources, err := setup.LoadResources()
	if err != nil {
		fmt.Println(err)
		return
	}
	apicall.Resources = resources

	apicaller := apicall.Caller{
		Getter: apicall.NewHttpGetter(),
	}

	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		queryParams := map[string]string{}
		c.Request().URI().QueryArgs().VisitAll(func(k, v []byte) {
			queryParams[string(k)] = string(v)
		})
		headers := map[string]string{}
		c.Request().Header.VisitAll(func(k, v []byte) {
			headers[string(k)] = string(v)
		})
		data, err := apicaller.Get("swapi-people", queryParams, headers)

		if err != nil {
			fmt.Println(err)
			return c.SendString("Error")
		}

		return c.JSON(data)
	})

	log.Fatal(app.Listen(":3000"))

}
