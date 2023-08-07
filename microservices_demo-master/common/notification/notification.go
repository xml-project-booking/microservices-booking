package notification

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Title      string
	Content    string
	Type       Type
	NotifierId primitive.ObjectID
	ActorId    primitive.ObjectID
}

type Type int8

const (
	ReservationRequestCreated Type = iota
	ReservationCancelled
	HostRated
	AccommodationRated
	HostFeaturedStatusChanged
	ReservationRequestResponded
)
