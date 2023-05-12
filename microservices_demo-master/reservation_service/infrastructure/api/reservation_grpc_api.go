package api

import (
	"context"
	pb "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	"resevation/application"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationHandler struct {
	pb.UnsafeReservationServiceServer
	service *application.ReservationService
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
