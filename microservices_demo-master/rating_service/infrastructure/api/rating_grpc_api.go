package api

import (
	"context"
	"fmt"
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

func (handler *RatingHandler) CanUserLeaveRating(ctx context.Context, request *pb.CanUserLeaveRatingRequest) (*pb.CanUserLeaveRatingResponse, error) {
	fmt.Println(request)
	canLeaveAccommodation := true
	canLeaveHost := true
	fmt.Println("usaoooo")
	ratings, err := handler.service.GetAll()
	for _, Rating := range ratings {
		if Rating.UserID.Hex() == request.UserId && Rating.TargetType == 0 && Rating.TargetId.Hex() == request.AccommodationId {
			canLeaveAccommodation = false
		}
	}
	for _, Rating := range ratings {
		if Rating.UserID.Hex() == request.UserId && Rating.TargetType == 1 && Rating.TargetId.Hex() == request.HostId {
			canLeaveHost = false
		}
	}
	if err != nil {
		return nil, err
	}
	return &pb.CanUserLeaveRatingResponse{
		CanLeaveAccommodation: canLeaveAccommodation,
		CanLeaveHost:          canLeaveHost,
	}, nil
}

func (handler *RatingHandler) CreateAccommodationRating(ctx context.Context, request *pb.CreateAccommodationRatingRequest) (*pb.CreateAccommodationRatingResponse, error) {
	fmt.Println(request)
	fmt.Println(request.Rating)
	rating := mapNewRating(request.Rating)
	err := handler.service.CreateRating(rating)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccommodationRatingResponse{
		Id: "kreirano",
	}, nil
}

func (handler *RatingHandler) CreateHostRating(ctx context.Context, request *pb.CreateHostRatingRequest) (*pb.CreateHostRatingResponse, error) {
	rating := mapNewRating(request.Rating)
	err := handler.service.CreateRating(rating)
	if err != nil {
		return nil, err
	}

	return &pb.CreateHostRatingResponse{
		Id: "kreirano",
	}, nil
}
func (handler *RatingHandler) GetAverageHostRating(ctx context.Context, request *pb.AverageHostRequest) (*pb.AverageHostResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	averageRating, _ := handler.service.GetHostAverage(objectId)
	response := &pb.AverageHostResponse{Average: float32(averageRating)}
	return response, nil
}
func (handler *RatingHandler) GetAverageAccommodationRating(ctx context.Context, request *pb.AverageHostRequest) (*pb.AverageHostResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	fmt.Println("ov je jebeni id")
	fmt.Println(objectId)
	averageRating, _ := handler.service.GetAccommodationAverage(objectId)
	response := &pb.AverageHostResponse{Average: float32(averageRating)}
	return response, nil
}
func (handler *RatingHandler) DeleteRating(ctx context.Context, request *pb.DeleteRatingRequest) (*pb.DeleteRatingResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Rating, err := handler.service.Get(objectId)
	boolVar := handler.service.Delete(Rating)
	fmt.Println(boolVar)
	if err != nil {
		return nil, err
	}

	//ReservationPb := mapReservation(Reservation)
	response := &pb.DeleteRatingResponse{
		Id: "nista",
	}
	return response, nil
}
func (handler *RatingHandler) GetRatingsByType(ctx context.Context, request *pb.GetRatingsByTypeRequest) (*pb.GetRatingsByTypeResponse, error) {
	id := request.Id
	targetType := request.Type
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ratings, err := handler.service.GetTargetRatings(objectId, targetType)

	if err != nil {
		return nil, err
	}
	response := &pb.GetRatingsByTypeResponse{
		Ratings: []*pb.Rating{},
	}
	for _, r := range ratings {
		current := mapRatingToPb(r)
		response.Ratings = append(response.Ratings, current)
	}

	return response, nil
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
	response := &pb.GetResponse{Rating: mapRatingToPb(Rating)}

	return response, nil
}

func (handler *RatingHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	_, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Ratings: []*pb.Rating{},
	}

	return response, nil
}
