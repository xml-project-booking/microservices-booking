package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	GetAll() ([]*Reservation, error)
	Insert(user *Reservation) error
	DeleteAll()
	UpdateStatus(user *Reservation) error
}
