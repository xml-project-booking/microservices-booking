package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	GetAll() ([]*Reservation, error)
	Insert(reservation *Reservation) error
	DeleteAll()
	UpdateStatusForCanceled(reservation *Reservation) error
	UpdateStatusForCanceledUser(reservation *Reservation) error
	UpdateStatusForConfirmed(reservation *Reservation) error
	GetAllReservationRequests() ([]*Reservation, error)
	GetAllReservation() ([]*Reservation, error)
	GetGuestAccommodationReservation(accommodationId, guestId primitive.ObjectID) ([]*Reservation, error)
	GetGuestAccommodationHostReservation(hostId, guestId primitive.ObjectID) ([]*Reservation, error)
	GetAllGuestReservation(guestId primitive.ObjectID) ([]*Reservation, error)
	GetAllReservationByAccommodation(accommodationId primitive.ObjectID) ([]*Reservation, error)
	GetAllConfirmedReservationByAccommodation(accommodationId primitive.ObjectID) ([]*Reservation, error)
	GetAllReservationByGuest(guestId primitive.ObjectID) ([]*Reservation, error)
	GetAllReservationByGuestPending(guestId primitive.ObjectID) ([]*Reservation, error)
	DeleteReservationById(reservationId primitive.ObjectID) bool
	GetReservationsByHost(hostId primitive.ObjectID) ([]*Reservation, error)
	GetReservationCancelByHost(hostId primitive.ObjectID) ([]*Reservation, error)
}
