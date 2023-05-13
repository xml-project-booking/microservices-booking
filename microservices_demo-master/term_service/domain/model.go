package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Term struct {
	Id primitive.ObjectID `bson:"_id"`
}
