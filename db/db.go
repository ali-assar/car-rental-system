package db

import "os"

var DBNAME = os.Getenv("MONGO_DB_NAME")

type Pagination struct {
	Limit int64
	Page  int64
}

type Store struct {
	User        UserStore
	Agency      AgencyStore
	Car         CarStore
	Reservation ReservationStore
}
