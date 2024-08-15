package handlers

import (
	"backend/database"

	"github.com/gofiber/fiber/v2"
)

func DeleteUserHandler(ctx *fiber.Ctx) error {
	name := ctx.Query("name")

	err := database.DeleteUser(name)

	if err != nil {
		return ctx.JSON(
			fiber.Map{
				"error": err,
				"body":  "failed to delete user...",
			},
		)
	}

	return ctx.JSON(
		fiber.Map{
			"error": nil,
			"body":  "successfully deleted user...",
		},
	)

}
