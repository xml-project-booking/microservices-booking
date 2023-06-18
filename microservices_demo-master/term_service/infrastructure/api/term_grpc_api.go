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

func (handler *TermHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {

	//OVO NZM STA JE VRV JSON U BINARNO
	jsonBytes, err := protojson.Marshal(request)
	if err != nil {
		{
			handler.LogError.WithFields(logrus.Fields{
				"status":    "failure",
				"location":  "Term Handler",
				"action":    "CRADA731",
				"timestamp": time.Now().String(),
			}).Error("Wrong cast json to Term!")
		}
	}

	err = json.Unmarshal(jsonBytes, &request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//KONVERZIJA IZ STRINGA U ODGOVARAJUCE TIPOVE
	layout := "2006-01-02T15:04:05.000Z"
	id, _ := primitive.ObjectIDFromHex(request.Id)
	accId, _ := primitive.ObjectIDFromHex(request.AccommodationId)
	usrId, _ := primitive.ObjectIDFromHex(request.UserId)

	date, _ := time.Parse(layout, request.Date)

	//PROVJER

	newTerm := domain.NewTermWithId(id, accId, usrId, request.PriceType, int32(request.Value), date)
	err = handler.service.Update(newTerm)
	if err != nil {
		return nil, err
	}
	res := pb.UpdateResponse{}

	return &res, nil
}

func (handler *TermHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	//OVO NZM STA JE VRV JSON U BINARNO
	jsonBytes, err := protojson.Marshal(request)
	if err != nil {
		{
			handler.LogError.WithFields(logrus.Fields{
				"status":    "failure",
				"location":  "Term Handler",
				"action":    "CRADA731",
				"timestamp": time.Now().String(),
			}).Error("Wrong cast json to Term!")
		}
	}

	err = json.Unmarshal(jsonBytes, &request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//KONVERZIJA IZ STRINGA U ODGOVARAJUCE TIPOVE

	id, _ := primitive.ObjectIDFromHex(request.Id)

	//PROVJER

	Term, err := handler.service.Get(id)
	err = handler.service.Delete(Term)

	if err != nil {
		return nil, err
	}
	res := pb.DeleteResponse{}

	return &res, nil
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

func (handler *TermHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.GetAllResponse, error) {

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
	layout := "2006-01-02T15:04:05.000Z"
	accId, _ := primitive.ObjectIDFromHex(request.AccommodationId)
	usrId, _ := primitive.ObjectIDFromHex(request.UserId)

	strtDate, _ := time.Parse(layout, request.StartDate)
	eDate, _ := time.Parse(layout, request.EndDate)

	//PROVJERE

	var isTaken = handler.service.CheckForReservationInDateRange(accId, strtDate, eDate)

	if isTaken {
		return nil, fmt.Errorf("reservation is already taken for the specified date range")
	}

	var Terms []*domain.Term // Ovo su termini koje smo napravili

	// Iteriranje kroz datume i pravljenje slobodnih
	for date := strtDate; date.Before(eDate); date = date.AddDate(0, 0, 1) {
		newTerm := domain.NewTerm(accId, usrId, request.PriceType, int32(request.Value), date)
		err = handler.service.Create(newTerm)
		Terms = append(Terms, newTerm)

		if err != nil {
			return nil, err
		}
	}

	// Konverzija i priprema za odgovor
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
