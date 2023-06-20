package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	Id                      primitive.ObjectID `bson:"_id"`
	Name                    string             `bson:"name"`
	ReservationConfirmation string             `bson:"reservation_confirmation"`
	City                    string             `bson:"city"`
	Country                 string             `bson:"country"`
	Street                  string             `bson:"street"`
	StreetNumber            string             `bson:"street_number"`
	MinGuest                int                `bson:"min_guest"`
	MaxGuest                int                `bson:"max_guest"`
	HostId                  primitive.ObjectID `bson:"host_id"`
	Wifi                    bool               `bson:"wifi"`
	Kitchen                 bool               `bson:"kitchen"`
	AirConditioning         bool               `bson:"air_conditioning"`
	FreeParking             bool               `bson:"free_parking"`
	Photos                  string             `bson:"photos"`
}

type ConfirmationType int

const (
	Automatic ConfirmationType = iota
	Manually
)

type ReservationDTO struct {
	Id              primitive.ObjectID
	AccommodationID primitive.ObjectID
	StartDate       string
	EndDate         string
	GuestNumber     string
	Confirmation    string
	GuestId         primitive.ObjectID
}

type AccommodationDTO struct {
	Id                      primitive.ObjectID
	Name                    string
	City                    string
	Country                 string
	Street                  string
	StreetNumber            string
	Amenities               []string
	MinGuest                string
	MaxGuest                string
	ReservationConfirmation string
}

type Address struct {
	Id           primitive.ObjectID
	Street       string
	StreetNumber string
	City         string
	Country      string
}
