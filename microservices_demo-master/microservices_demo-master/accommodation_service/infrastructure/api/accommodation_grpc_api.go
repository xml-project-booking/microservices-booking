package api

import (
	"accommodation_service/application"
	"accommodation_service/domain"
	"context"
	pb "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Accommodation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	AccommodationPb := mapAccommodation(Accommodation)
	response := &pb.GetResponse{
		Accommodation: AccommodationPb,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Accommodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, Accommodation := range Accommodations {
		current := mapAccommodation(Accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAllIdsByHost(ctx context.Context, request *pb.GetAllIdsByHostRequest) (*pb.GetAllIdsByHostResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	allAccommodations, err2 := handler.service.GetAll()
	if err2 != nil {
		return nil, err2
	}
	accommodations := []*domain.Accommodation{}
	for _, a := range allAccommodations {
		if a.HostId == objectId {
			accommodations = append(accommodations, a)
		}
	}

	accommodationIds := []string{}
	for _, accommodation := range accommodations {
		accommodationIds = append(accommodationIds, accommodation.HostId.Hex())
	}
	response := &pb.GetAllIdsByHostResponse{
		Ids: accommodationIds,
	}
	return response, nil
}
func (handler *AccommodationHandler) DeleteAllByHost(ctx context.Context, request *pb.DeleteAllByHostRequest) (*pb.DeleteAllByHostResponse, error) {
	err := handler.service.DeleteAllAccommodationsByHost(request.Id)

	if err != nil {
		return &pb.DeleteAllByHostResponse{
			RequestResult: &pb.RequestResult{
				Code:    400,
				Message: err.Error(),
			},
		}, err
	}

	return &pb.DeleteAllByHostResponse{
		RequestResult: &pb.RequestResult{
			Code: 200,
		},
	}, nil
}
