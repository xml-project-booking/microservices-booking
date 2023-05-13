package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Term struct {
	Id              primitive.ObjectID `bson:"id"`              // ID termina
	AccommodationID primitive.ObjectID `bson:"accommodationId"` // ID smeštaja za koji je termin vezan
	UserID          primitive.ObjectID `bson:"userId"`          // ID korisnika koji je rezervisao smeštaj
	Date            time.Time          `bson:"date"`            // Datum termina
	Available       bool               `bson:"available"`       // Dostupnost smeštaja za dati termin
	Price           float64            `bson:"price"`           // Cena smeštaja za dati termin
	//PriceType       string             // Tip cene - "per_guest" ili "per_unit" dodati na nivou nekretnine
}

// Metoda za kreiranje novog termina
func NewTerm(accommodationID, id primitive.ObjectID, date time.Time, available bool, price float64, userID primitive.ObjectID) *Term {
	return &Term{
		AccommodationID: accommodationID,
		Id:              id,
		Date:            date,
		Available:       available,
		Price:           price,
		UserID:          userID,
	}
}

// Metoda za izmenu cene termina
func (t *Term) UpdatePrice(newPrice float64) {
	t.Price = newPrice
}

// Metoda za proveru da li je termin zauzet
func (t *Term) IsReserved() bool {
	return t.UserID != primitive.NilObjectID
}
