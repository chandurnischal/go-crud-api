package handlers

import (
	"backend/database"

	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(ctx *fiber.Ctx) error {
	name, location := ctx.Query("name"), ctx.Query("location")

	if name == "" && location == "" {
		return ctx.JSON(
			fiber.Map{
				"error": "unable to create user with empty name or location",
				"body":  nil,
			},
		)
	}

	err := database.CreateUser(name, location)

	if err != nil {
		return ctx.JSON(
			fiber.Map{
				"error": err,
				"body":  "failed to create user",
			},
		)
	}
	return ctx.JSON(
		fiber.Map{
			"error": nil,
			"body":  "successfully created user",
		},
	)
}
