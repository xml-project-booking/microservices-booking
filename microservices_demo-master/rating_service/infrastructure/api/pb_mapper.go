package api

import (
	"fmt"
	pb "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_service/domain"
	"time"
)

func mapNewRating(ratingPb *pb.Rating) *domain.Rating {
	layout := "2006-01-02T15:04:05.000Z"
	parsedTime, err := time.Parse(layout, ratingPb.LastModified)
	if err != nil {
		fmt.Println("Error:", err)

	}
	rating := &domain.Rating{
		UserID:       getObjectId(ratingPb.UserID),
		TargetId:     getObjectId(ratingPb.TargetId),
		RatingValue:  ratingPb.RatingValue,
		LastModified: parsedTime,
		TargetType:   int(ratingPb.TargetType),
	}
	return rating
}

func mapRatingToPb(rating *domain.Rating) *pb.Rating {
	formattedTime := rating.LastModified.Format("2006-01-02T15:04:05Z")
	ratingPb := &pb.Rating{
		UserID:       rating.UserID.Hex(),
		TargetId:     rating.TargetId.Hex(),
		RatingValue:  rating.RatingValue,
		LastModified: formattedTime,
		Id:           rating.Id.Hex(),
	}
	return ratingPb
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
