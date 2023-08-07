package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Rating struct {
	Id           primitive.ObjectID `bson:"_id"` // ID smeštaja za koji je termin vezan
	UserID       primitive.ObjectID `bson:"userId"`
	RatingValue  int32
	TargetType   int `bson:"target_type"`
	LastModified time.Time
	TargetId     primitive.ObjectID `bson:"targetId"`
	RatingStatus string             // ID korisnika koji je rezervisao smeštaj

}
type RatingDTO struct {
	Id           primitive.ObjectID // ID smeštaja za koji je termin vezan
	UserID       primitive.ObjectID
	RatingValue  int32
	TargetType   int
	LastModified time.Time
	TargetId     primitive.ObjectID
}

// Metoda za proveru da li je termin zauzet
