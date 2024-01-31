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
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
