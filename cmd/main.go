package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/thslopes/bff/apicall"
	"github.com/thslopes/bff/setup"

	"github.com/gofiber/fiber/v3"
)

var apicaller = apicall.Caller{
	Getter: apicall.NewHttpGetter(),
}

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

	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/:query", handler)
	app.Post("/:query", handler)

	log.Fatal(app.Listen(":3000"))

}

func handler(c fiber.Ctx) error {
	queryParams := map[string]string{}
	c.Request().URI().QueryArgs().VisitAll(func(k, v []byte) {
		queryParams[string(k)] = string(v)
	})
	headers := map[string]string{}
	c.Request().Header.VisitAll(func(k, v []byte) {
		headers[string(k)] = string(v)
	})

	body := c.Request().Body()
	var bodyData interface{}
	err := json.Unmarshal(body, &bodyData)
	if err != nil {
		return err
	}

	data, err := apicaller.Do(c.Params("query"), queryParams, headers, bodyData)

	if err != nil {
		return err
	}

	return c.JSON(data)
}
