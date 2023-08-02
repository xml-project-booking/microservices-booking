package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	pb "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/encoding/protojson"
	"rating_service/application"
	"time"
)

type RatingHandler struct {
	pb.UnimplementedTermServiceServer
	service  *application.RatingService
	LogInfo  *logrus.Logger
	LogError *logrus.Logger
}

func NewUserHandler(service *application.TermService) *RatingHandler {
	return &RatingHandler{
		service: service,
	}
}

func (handler *RatingHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	//OVO NZM STA JE VRV JSON U BINARNO

}

func (handler *RatingHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Term, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	response := mapTerm(Term)
	return response, nil
}

func (handler *RatingHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Terms, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		GetResponses: []*pb.GetResponse{},
	}
	for _, Term := range Terms {
		current := mapTerm(Term)
		response.GetResponses = append(response.GetResponses, current)
	}
	return response, nil
}

func (handler *RatingHandler) GetByAccommodationId(ctx context.Context, request *pb.GetByAccommodationIdRequest) (*pb.GetByAccommodationIdResponse, error) {
	accId, _ := primitive.ObjectIDFromHex(request.AccommodationId)
	Terms, err := handler.service.GetByAccommodationId(accId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetByAccommodationIdResponse{
		GetResponses: []*pb.GetResponse{},
	}
	for _, Term := range Terms {
		current := mapTerm(Term)
		response.GetResponses = append(response.GetResponses, current)
	}
	return response, nil

}

func (handler *RatingHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.GetAllResponse, error) {

	//OVO NZM STA JE VRV JSON U BINARNO
	jsonBytes, err := protojson.Marshal(request)
	if err != nil {
		{
			handler.LogError.WithFields(logrus.Fields{
				"status":    "failure",
				"location":  "Reservation Handler",
				"action":    "CRADA731",
				"timestamp": time.Now().String(),
			}).Error("Wrong cast json to ReservationDTO!")
		}
	}

	err = json.Unmarshal(jsonBytes, &request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//KONVERZIJA IZ STRINGA U ODGOVARAJUCE TIPOVE

}
