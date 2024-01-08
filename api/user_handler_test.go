// To do make the user test better and also test put and delete

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/Ali-Assar/car-rental-system/db"
	"github.com/Ali-Assar/car-rental-system/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDbUri = "mongodb://localhost:27017"

//const dbname = "car-rental-system-test"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) tearDown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDbUri))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func createUserAndPost(t *testing.T, app *fiber.App, userHandler *UserHandler, params types.CreateUserParams) *types.User {
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	return &user
}

func TestPostUser(t *testing.T) {
	tdb := setup(t) //
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser) //

	params := types.CreateUserParams{
		Email:     "chifo@kongfo.com",
		FirstName: "chifo",
		LastName:  "master",
		Password:  "12345678",
	}
	user := createUserAndPost(t, app, userHandler, params)

	if len(user.ID) == 0 {
		t.Errorf("expecting a user id")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("EncryptedPassword should not included in json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected the first name be %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected the last name be %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected the email be %s but got %s", params.Email, user.Email)
	}

}

func TestGetUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	user1Params := &types.CreateUserParams{
		Email:     "user1@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password123",
	}
	user2Params := &types.CreateUserParams{
		Email:     "user2@example.com",
		FirstName: "Jane",
		LastName:  "Doe",
		Password:  "password456",
	}

	createUserAndPost(t, app, userHandler, *user1Params)
	createUserAndPost(t, app, userHandler, *user2Params)

	app.Get("/", userHandler.HandleGetUsers)

	req := httptest.NewRequest("GET", "/", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, but got %d", resp.StatusCode)
	}

	var users []*types.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, but got %d", len(users))
	}
	if users[0].FirstName != "John" {
		t.Errorf("expected the first user's first name to be John, but got %s", users[0].FirstName)
	}
	if users[1].LastName != "Doe" {
		t.Errorf("expected the second user's last name to be Doe, but got %s", users[1].LastName)
	}
}

func TestGetUserByID(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := &types.CreateUserParams{
		Email:     "user1@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password123",
	}
	user := createUserAndPost(t, app, userHandler, *params)

	app.Get("/:id", userHandler.HandleGetUser)

	oid, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/"+oid.Hex(), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, but got %d", resp.StatusCode)
	}

	var retrievedUser types.User
	err = json.NewDecoder(resp.Body).Decode(&retrievedUser)
	if err != nil {
		t.Fatal(err)
	}
	if retrievedUser.ID != user.ID {
		t.Errorf("expected the retrieved user ID to be %s, but got %s", user.ID, retrievedUser.ID)
	}
	if retrievedUser.FirstName != user.FirstName {
		t.Errorf("expected the retrieved user's first name to be %s, but got %s", user.FirstName, retrievedUser.FirstName)
	}
	if retrievedUser.LastName != user.LastName {
		t.Errorf("expected the retrieved user's last name to be %s, but got %s", user.LastName, retrievedUser.LastName)
	}
}

func TestDeleteUserByID(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := &types.CreateUserParams{
		Email:     "user1@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password123",
	}
	user := createUserAndPost(t, app, userHandler, *params)
	app.Delete("/:id", userHandler.HandleDeleteUser)

	oid, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("DELETE", "/"+oid.Hex(), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, but got %d", resp.StatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := &types.CreateUserParams{
		Email:     "user1@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password123",
	}
	user := createUserAndPost(t, app, userHandler, *params)
	app.Put("/:id", userHandler.HandlePutUser)

	updatedParams := &types.CreateUserParams{
		FirstName: "poo",
		LastName:  "dragonWarrior",
	}

	oid, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(updatedParams)
	req := httptest.NewRequest("PUT", "/"+oid.Hex(), bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	var updatedUser types.User
	json.NewDecoder(resp.Body).Decode(&updatedUser)
	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, but got %d", resp.StatusCode)
	}
	if updatedUser.FirstName == user.FirstName {
		t.Errorf("expected the first name to be updated, but it's still %s", user.FirstName)
	}
	if updatedUser.LastName == user.LastName {
		t.Errorf("expected the last name to be updated, but it's still %s", user.LastName)
	}
}
