package main

import (
	"context"
	"flag"

	"log"

	"github.com/Ali-Assar/car-rental-system/api"
	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create a new fiber instance with custom config
var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "this is the default address API server") // you can change the port by this command  ./bin/api --listenAddr :7000
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	//handlers initialization
	var (
		userHandler   = api.NewUserHandler(db.NewMongoUserStore(client, db.DBNAME))
		agencyStore   = db.NewMongoAgencyStore(client)
		carStore      = db.NewMongoCarStore(client, agencyStore)
		agencyHandler = api.NewAgencyHandler(agencyStore, carStore)

		app   = fiber.New(config)
		apiv1 = app.Group("/api/v1")
	)

	//user handlers
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user/", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	//agency handlers
	apiv1.Get("/agency", agencyHandler.HandleGetAgencies)

	app.Listen(*listenAddr)
}
