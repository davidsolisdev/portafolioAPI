package main

import (
	"github.com/davidsolisdev/portafolioApi/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func App() (app *fiber.App) {
	app = fiber.New(fiber.Config{Prefork: true})

	app.Use(cors.New(cors.Config{}))

	app.Use(recover.New())

	app.Post("/contact-me", func(ctx *fiber.Ctx) error {
		var body *BodySendEmail = new(BodySendEmail)

		err := ctx.BodyParser(body)
		if err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		_, errr := utils.SendEmail(&utils.NewEmail{From: body.Email, To: "davidsolisdev@gmail.com", Subject: body.Email + ": " + body.Name}, body.Message)
		if errr != nil {
			return ctx.Status(400).SendString(errr.Error())
		}

		return ctx.SendStatus(200)
	})

	return app
}

type BodySendEmail struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
