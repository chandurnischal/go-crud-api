package handlers

import (
	"backend/database"

	"github.com/gofiber/fiber/v2"
)

func RetrieveByNameHandler(ctx *fiber.Ctx) error {
	name := ctx.Query("name")

	if name == "" {
		return RetrieveAllHandler(ctx)
	}

	user, err := database.RetrieveByName(name)

	if err != nil {
		return ctx.JSON(
			fiber.Map{
				"body":  nil,
				"error": "no user found by that name",
			},
		)
	}

	return ctx.JSON(
		fiber.Map{
			"error": nil,
			"body":  user,
		},
	)

}

func RetrieveAllHandler(ctx *fiber.Ctx) error {
	users, err := database.RetrieveAll()

	if err != nil {
		return ctx.JSON(
			fiber.Map{
				"body":  "failed to retrieve users...",
				"error": err,
			},
		)
	}

	return ctx.JSON(
		fiber.Map{
			"body":  users,
			"error": nil,
		},
	)
}
