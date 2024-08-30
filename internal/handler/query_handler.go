package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/thslopes/bff/internal/request"
)

func QueryHandler(c fiber.Ctx) error {
	req := request.NewRequest(c.Params("query"))

	c.Request().URI().QueryArgs().VisitAll(func(k, v []byte) {
		req.QueryParams[string(k)] = string(v)
	})

	c.Request().Header.VisitAll(func(k, v []byte) {
		req.Headers[string(k)] = string(v)
	})

	body := c.Request().Body()
	if len(body) > 0 {
		err := json.Unmarshal(body, &req.Body)
		if err != nil {
			return err
		}
	}

	data, err := req.Execute()

	if err != nil {
		return err
	}

	return c.JSON(data)
}
