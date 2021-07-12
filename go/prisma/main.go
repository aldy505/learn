package main

import (
	"context"
	"log"

	"github.com/aldy505/learn-go/prisma/db"
	"github.com/aldy505/learn-go/prisma/model"
	phccrypto "github.com/aldy505/phc-crypto"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	crypto, err := phccrypto.Use("scrypt", phccrypto.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	client := db.NewClient()

	err = client.Prisma.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	app.Get("/user", func(c *fiber.Ctx) error {
		users, err := client.User.FindMany().Exec(context.Background())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"users": users,
			"total": len(users),
		})
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		var user model.User
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		password, err := crypto.Hash(user.Password)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		createdUser, err := client.User.CreateOne(
			db.User.Name.Set(user.Name),
			db.User.Email.Set(user.Email),
			db.User.Password.Set(password),
		).Exec(context.Background())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"user": createdUser,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
