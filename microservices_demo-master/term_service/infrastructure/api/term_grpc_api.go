package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	pb "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/encoding/protojson"
	"term_service/application"
	"term_service/domain"
	"time"
)

type TermHandler struct {
	pb.UnimplementedTermServiceServer
	service  *application.TermService
	LogInfo  *logrus.Logger
	LogError *logrus.Logger
}

func NewUserHandler(service *application.TermService) *TermHandler {
	return &TermHandler{
		service: service,
	}
}

func (handler *TermHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
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

func (handler *TermHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
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

func (handler *TermHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {

	fmt.Println(request.String())
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

	layout := "2006-01-02T15:04:05.000Z"
	accId, _ := primitive.ObjectIDFromHex(request.AccommodationId)
	usrId, _ := primitive.ObjectIDFromHex(request.UserId)

	strtDate, _ := time.Parse(layout, request.StartDate)
	eDate, _ := time.Parse(layout, request.EndDate)

	newTerm := domain.NewTerm(accId, usrId, request.PriceType, int32(request.Value), strtDate, eDate)

	err = handler.service.Create(newTerm)
	TermPb := mapTerm(newTerm)
	fmt.Println(TermPb.Id)

	if err != nil {
		return nil, err
	}
	response := &pb.CreateResponse{
		Id: TermPb.Id,
	}

	return response, nil
}
