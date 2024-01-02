package main

import (
	"flag"

	"github.com/Ali-Assar/reservation-system/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "this is the default address API server") // you can change the port by this command  ./bin/api --listenAddr :7000
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user", api.HandleGetUser)
	app.Get("/poo", handlePoo)

	app.Listen(*listenAddr)
}

func handlePoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "working just fine"})
}
