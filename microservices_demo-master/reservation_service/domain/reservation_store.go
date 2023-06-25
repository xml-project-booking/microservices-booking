package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	GetAll() ([]*Reservation, error)
	Insert(reservation *Reservation) error
	DeleteAll()
	UpdateStatusForCanceled(reservation *Reservation) error
	UpdateStatusForConfirmed(reservation *Reservation) error
	GetAllReservationRequests() ([]*Reservation, error)
	GetAllReservation() ([]*Reservation, error)
	GetAllGuestReservation(guestId primitive.ObjectID) ([]*Reservation, error)
	GetAllReservationByAccommodation(accommodationId primitive.ObjectID) ([]*Reservation, error)
	DeleteReservationById(reservationId primitive.ObjectID) bool
}
