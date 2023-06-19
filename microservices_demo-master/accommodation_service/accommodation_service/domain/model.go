package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	HostId primitive.ObjectID `bson:"host_id"`
}
