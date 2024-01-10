package db

const (
	DBNAME     = "car-rental-system"
	TestDBNAME = "car-rental-system-Test"
	DBURI      = "mongodb://localhost:27017"
)

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
