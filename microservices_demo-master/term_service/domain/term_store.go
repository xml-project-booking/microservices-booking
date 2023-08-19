package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TermStore interface {
	Get(id primitive.ObjectID) (*Term, error)
	GetAll() ([]*Term, error)
	Insert(term *Term) error
	Delete(term *Term) error
	DeleteAll()
	Update(term *Term) error
	GetByAccommodationId(id primitive.ObjectID) ([]*Term, error)
	GetByAccommodationIdOne(id primitive.ObjectID) (*Term, error)
	GetTermsInPriceRange(minPrice, maxPrice int32) ([]*Term, error)
}
