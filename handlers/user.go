package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vovah1a/myMoney_v2/database"
	"github.com/vovah1a/myMoney_v2/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Div Vova Trivia App!")
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}
