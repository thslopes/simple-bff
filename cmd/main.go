package main

import (
	"fmt"
	"log"

	"github.com/thslopes/bff/apicall"
	"github.com/thslopes/bff/setup"

	"github.com/gofiber/fiber/v3"
)

func main() {
	mappings, err := setup.LoadApiCallFromFile("./mappings/query.json")
	if err != nil {
		fmt.Println(err)
		return
	}

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
		for _, v := range mappings {
			data, err := apicaller.Get(v, queryParams, headers)

			if err != nil {
				fmt.Println(err)
				return c.SendString("Error")
			}

			return c.SendString(string(data))
		}
		return c.SendString("no")
	})

	log.Fatal(app.Listen(":3000"))

}
