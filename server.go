package main

import (
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	ApiKey    string `json:"api_key"`
	Team1Name string `json:"team_1_name"`
	Team1Yr   string `json:"team_1_yr"`
	Team2Name string `json:"team_2_name"`
	Team2Yr   string `json:"team_2_yr"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		req := new(Request)
		err := c.BodyParser(req)
		if err != nil {
			return c.SendStatus(400)
		}

		return c.JSON(req)
	})

	app.Listen(":8080")
}
