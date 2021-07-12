package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/", getAllUser)
	user.Get("/:username", getByUserName)
	user.Post("/", addNewUser)
	user.Patch("/:user", editUser)
	user.Put("/:username", replaceUser)
	user.Delete("/:username", deleteUser)
}
