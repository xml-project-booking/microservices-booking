package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Term struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationID primitive.ObjectID `bson:"accommodationId"` // ID smeštaja za koji je termin vezan
	UserID          primitive.ObjectID `bson:"userId"`          // ID korisnika koji je rezervisao smeštaj
	PriceType       string             `bson:"price_type"`
	Value           int32              `bson:"value"`
	Date            time.Time          `bson:"date"`
}

func NewTerm(accommodationID primitive.ObjectID, userID primitive.ObjectID, priceType string, value int32, date time.Time) *Term {
	return &Term{AccommodationID: accommodationID, UserID: userID, PriceType: priceType, Value: value, Date: date}
}

// Metoda za proveru da li je termin zauzet
func (t *Term) IsReserved() bool {
	return t.UserID != primitive.NilObjectID
}
