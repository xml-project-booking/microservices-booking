package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingStore interface {
	Get(id primitive.ObjectID) (*Rating, error)
	GetAll() ([]*Rating, error)
	Insert(term *Rating) error
	Delete(term *Rating) error
	DeleteAll()
	Update(term *Rating) error
	GetByAccommodationId(id primitive.ObjectID) ([]*Rating, error)
}
