package main

import (
	"context"
	"flag"

	"log"

	"github.com/Ali-Assar/car-rental-system/api"
	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/middleware"
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
		agencyStore      = db.NewMongoAgencyStore(client)
		carStore         = db.NewMongoCarStore(client, agencyStore)
		userStore        = db.NewMongoUserStore(client)
		reservationStore = db.NewMongoReservationStore(client)
		store            = &db.Store{
			Agency:      agencyStore,
			Car:         carStore,
			User:        userStore,
			Reservation: reservationStore,
		}

		userHandler   = api.NewUserHandler(userStore)
		agencyHandler = api.NewAgencyHandler(store)
		authHandler   = api.NewAuthHandler(userStore)
		carHandler    = api.NewCarHandler(store)
		app           = fiber.New(config)
		auth          = app.Group("/api")
		apiv1         = app.Group("/api/v1", middleware.JWTAuthentication(userStore))
	)

	//auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	//user handlers
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user/", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	//agency handlers
	apiv1.Get("/agency", agencyHandler.HandleGetAgencies)
	apiv1.Get("/agency/:id", agencyHandler.HandleGetAgency)
	apiv1.Get("/agency/:id/cars", agencyHandler.HandleGetCars)
	apiv1.Get("/car", carHandler.HandleGetCars)
	apiv1.Post("car/:id/reservation", carHandler.HandleReserveCar)

	app.Listen(*listenAddr)
}
