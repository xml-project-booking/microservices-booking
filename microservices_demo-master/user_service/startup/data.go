package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user_service/domain"
)

var users = []*domain.User{
	{
		Id:                 getObjectId("623b0cc336a1d6fd8c1cf0f6"),
		CancellationNumber: 1,
		Name:               "Marko",
		Surname:            "Markovic",
		Role:               1,
	},
	{
		Id:                 getObjectId("623b0cc336a1d6fd8c1cf0f5"),
		CancellationNumber: 1,
		Name:               "Ivan",
		Surname:            "Ivanovic",
		Role:               1,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
