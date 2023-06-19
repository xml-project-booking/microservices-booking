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
}
