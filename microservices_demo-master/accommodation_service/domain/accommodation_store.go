package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccommodationStore interface {
	Get(id primitive.ObjectID) (*Accommodation, error)
	GetAll() ([]*Accommodation, error)
	Insert(user *Accommodation) error
	UpdateReservationConfirmationType(accommodation *Accommodation) error
	DeleteAll()
	UpdateStatus(user *Accommodation) error
	DeleteAccommodation(id primitive.ObjectID) error
	FilterAccommodationsByAmenities(amenities []bool, accommodations []*Accommodation) ([]*Accommodation, error)
}
