package main

import (
	"context"
	"flag"
	"net/http"

	"log"

	"github.com/Ali-Assar/car-rental-system/api"
	"github.com/Ali-Assar/car-rental-system/db"

	//"github.com/Ali-Assar/car-rental-system/middleware"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create a new fiber instance with custom config
var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(api.Error); ok {
			return c.Status(apiError.Code).JSON(apiError)
		}
		apiError := api.NewError(http.StatusInternalServerError, err.Error())
		return c.Status(apiError.Code).JSON(apiError)
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

	app.Listen(*listenAddr)
}
