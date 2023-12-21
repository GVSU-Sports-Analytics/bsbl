package main

import (
	"github.com/GVSU-Sports-Analytics/bsbl-sim-tui/bsbl"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		req := new(bsbl.Request)
		err := c.BodyParser(req)
		if err != nil {
			return c.SendStatus(400)
		}

		_, _, err = bsbl.GetTeams(req)
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(req)
	})

	app.Listen(":8080")
}
