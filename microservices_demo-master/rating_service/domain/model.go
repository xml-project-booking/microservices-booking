package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationID primitive.ObjectID `bson:"accommodationId"` // ID smeštaja za koji je termin vezan
	UserID          primitive.ObjectID `bson:"userId"`          // ID korisnika koji je rezervisao smeštaj

}

// Metoda za proveru da li je termin zauzet
