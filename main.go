package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", handleUser)
	app.Get("/poo", handlePoo)

	app.Listen(":5000")
}

func handlePoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "working just fine"})
}
func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "ali"})
}
