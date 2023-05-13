package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TermStore interface {
	Get(id primitive.ObjectID) (*Term, error)
	GetAll() ([]*Term, error)
	Insert(user *Term) error
	DeleteAll()
	UpdateStatus(user *Term) error
}
