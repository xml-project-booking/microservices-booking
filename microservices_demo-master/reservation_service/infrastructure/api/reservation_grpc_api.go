package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	pb "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	"google.golang.org/protobuf/encoding/protojson"
	"strconv"

	_ "github.com/golang/protobuf/jsonpb"
	"resevation/domain"

	"resevation/application"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	service  *application.ReservationService
	LogInfo  *logrus.Logger
	LogError *logrus.Logger
}

func (handler *ReservationHandler) MakeRequestForReservation(ctx context.Context, request *pb.ReservationRequest) (*pb.ReservationRequestResponse, error) {

	var reservationDTO domain.ReservationDTO
	fmt.Print("request: ")
	fmt.Println(request)

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

	err = json.Unmarshal(jsonBytes, &reservationDTO)
	if err != nil {
		// Handle error
	}
	fmt.Println("kako se ispisati  resefvationdto")
	fmt.Println(reservationDTO)

	/*if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "AdvertisementHandler",
			"action":    "CRADA731",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to AdvertisementDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/
	layout := "2006-01-02T15:04:05.000Z"
	startDate, _ := time.Parse(layout, reservationDTO.StartDate)
	endDate, _ := time.Parse(layout, reservationDTO.EndDate)
	num, err := strconv.Atoi(reservationDTO.GuestNumber)
	b, err := strconv.ParseBool(reservationDTO.Confirmation)

	var isTaken = handler.service.CheckForReservationInDateRange(reservationDTO.AccommodationID, startDate, endDate)
	re := &pb.ReservationRequestResponse{Id: "greska"}
	if isTaken == true {
		return re, nil
	}

	reservationRequest := domain.Reservation{
		AccommodationID: reservationDTO.AccommodationID,
		StartDate:       startDate,
		EndDate:         endDate,
		Confirmation:    b,
		GuestNumber:     int64(num),
	}
	fmt.Println(reservationDTO.AccommodationID)
	fmt.Println("ahhahahahahahahaahahahah")
	fmt.Println(reservationDTO.AccommodationID)
	fmt.Println(reservationDTO.StartDate)
	fmt.Println(reservationDTO.EndDate)

	err = handler.service.CreateReservationRequest(&reservationRequest)
	ReservationPb := mapReservation(&reservationRequest)
	response := &pb.ReservationRequestResponse{
		Id: ReservationPb.Id,
	}
	return response, nil
}

func NewReservationHandler(service *application.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (handler *ReservationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Reservation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	ReservationPb := mapReservation(Reservation)
	response := &pb.GetResponse{
		Reservation: ReservationPb,
	}
	return response, nil
}

func (handler *ReservationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Reservations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, Reservation := range Reservations {
		current := mapReservation(Reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

/*if err != nil {
	handler.LogError.WithFields(logrus.Fields{
		"status":    "failure",
		"location":  "AdvertisementHandler",
		"action":    "CRADA731",
		"timestamp": time.Now().String(),
	}).Error("Failed creating reservation!")
	.WriteHeader(http.StatusExpectationFailed)
	return
}*/

/*handler.LogInfo.WithFields(logrus.Fields{
	"status":    "success",
	"location":  "ReservationHandler",
	"action":    "CRADA731",
	"timestamp": time.Now().String(),
}).Info("Successfully created reservation request!")
w.WriteHeader(http.StatusCreated)
w.Header().Set("Content-Type", "application/json")*/
