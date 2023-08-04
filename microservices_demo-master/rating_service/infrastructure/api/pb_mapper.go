package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_service/domain"
)

func mapNewRating(ratingPb *pb.Rating) *domain.Rating {
	rating := &domain.Rating{
		UserID:      getObjectId(ratingPb.UserID),
		TargetId:    getObjectId(ratingPb.TargetId),
		RatingValue: ratingPb.RatingValue,
	}
	return rating
}

func mapRatingToPb(rating *domain.Rating) *pb.Rating {
	ratingPb := &pb.Rating{
		UserID:      rating.UserID.Hex(),
		TargetId:    rating.TargetId.Hex(),
		RatingValue: rating.RatingValue,
	}
	return ratingPb
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
