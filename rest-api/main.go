package main

import (
	"context"
	"os"

	"log"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/rest-api/api"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create a new fiber instance with custom config
var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
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

		userHandler        = api.NewUserHandler(userStore)
		agencyHandler      = api.NewAgencyHandler(store)
		authHandler        = api.NewAuthHandler(userStore)
		carHandler         = api.NewCarHandler(store)
		reservationHandler = api.NewReservationHandler(store)
		app                = fiber.New(config)
		auth               = app.Group("/api")
		apiv1              = app.Group("/api/v1", api.JWTAuthentication(userStore))
		admin              = apiv1.Group("/admin", api.AdminAuth)
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

	//car handlers
	apiv1.Get("/car", carHandler.HandleGetCars)
	apiv1.Post("car/:id/reservation", carHandler.HandleReserveCar)

	//reservation handler
	//apiv1.Get("/reservation", reservationHandler.HandleGetReservations)
	apiv1.Get("/reservation/:id", reservationHandler.HandleGetReservation)
	apiv1.Get("reservation/:id/cancel", reservationHandler.HandleCancelReservation)

	//admin handlers
	admin.Get("/reservation", reservationHandler.HandleGetReservations)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}
