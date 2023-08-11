package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_service/domain"
)

var ratings = []*domain.Rating{
	{
		Id: getObjectId("623b0cc336a1d6fd8c1cf0f6"),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
