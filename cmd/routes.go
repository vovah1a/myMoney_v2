package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vovah1a/myMoney_v2/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/user", handlers.CreateUser)
}
