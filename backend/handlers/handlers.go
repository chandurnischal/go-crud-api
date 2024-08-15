package handlers

import (
	"backend/database"

	"github.com/gofiber/fiber/v2"
)

func CheckDatabaseConnection(ctx *fiber.Ctx) error {
	if err := database.DB.Ping(); err != nil {
		return ctx.JSON(
			fiber.Map{
				"error": err,
				"body":  "failed to ping database",
			},
		)
	}
	return ctx.JSON(
		fiber.Map{
			"error":   nil,
			"message": "connected to database",
		},
	)
}
