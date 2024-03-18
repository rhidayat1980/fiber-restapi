package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rhidayat1980/fiber-restapi/handlers"
	"github.com/rhidayat1980/fiber-restapi/models"
)

func Authenticated(c *fiber.Ctx) error {
	json := new(models.Session)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Session Format",
		})
	}

	user, err := handlers.GetUser(json.Sessionid)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "404: not found",
		})
	}
	c.Locals("user", user)
	return c.Next()
}
