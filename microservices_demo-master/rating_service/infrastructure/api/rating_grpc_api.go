package api

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"rating_service/application"
)

type RatingHandler struct {
	pb.UnimplementedRatingServiceServer
	service  *application.RatingService
	LogInfo  *logrus.Logger
	LogError *logrus.Logger
}

func NewUserHandler(service *application.RatingService) *RatingHandler {
	return &RatingHandler{
		service: service,
	}
}

func (handler *RatingHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Rating, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	response := mapRating(Rating)
	return response, nil
}

func (handler *RatingHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	_, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Orders: []*pb.Order{},
	}

	return response, nil
}
