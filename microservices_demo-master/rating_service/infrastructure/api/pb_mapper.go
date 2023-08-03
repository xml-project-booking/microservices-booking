package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	"rating_service/domain"
)

func mapRating(term *domain.Rating) *pb.GetResponse {
	ratingPb := &pb.GetResponse{}
	return ratingPb
}
