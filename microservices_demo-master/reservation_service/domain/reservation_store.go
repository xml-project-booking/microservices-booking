package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	GetAll() ([]*Reservation, error)
	Insert(reservation *Reservation) error
	DeleteAll()
	UpdateStatus(reservation *Reservation) error
	GetAllReservationRequests() ([]*Reservation, error)
	GetAllReservation() ([]*Reservation, error)
}
