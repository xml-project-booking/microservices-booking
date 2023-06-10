package domain

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type ReservationDTO struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationID primitive.ObjectID `bson:"accommodation_id"`
	StartDate       string             `bson:"start_date"`
	EndDate         string             `bson:"end_date"`
	GuestNumber     string             `bson:"guest_number"`
	Confirmation    string             `bson:"confirmation"`
}

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationID primitive.ObjectID `bson:"accommodation_id"`
	StartDate       time.Time          `bson:"start_date"`
	EndDate         time.Time          `bson:"end_date"`
	GuestNumber     int64              `bson:"guest_number"`
	Confirmation    bool               `bson:"confirmation"`
}

type AccommodationRequest struct {
}

func (u *ReservationDTO) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *ReservationDTO) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
