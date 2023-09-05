package api

import (
	"accommodation_service/application"
	"accommodation_service/domain"
	"context"
	"fmt"

	//"encoding/json"
	//"fmt"

	pb "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"google.golang.org/protobuf/encoding/protojson"
	"strconv"
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
func (handler *AccommodationHandler) SearchAccommodationsByLocation(ctx context.Context, request *pb.SearchAccommodationRequest) (*pb.SearchAccommodationResponse, error) {
	accommodationsPb := request.Accommodations
	location := request.Location
	var accommodationsToSearch []*domain.Accommodation
	for _, Accommodation := range accommodationsPb {
		accommodationDomain := mapAccommodationPb(Accommodation)
		accommodationsToSearch = append(accommodationsToSearch, accommodationDomain)

	}
	accommodations := handler.service.SearchAccommodationsByLocation(accommodationsToSearch, location)
	response := &pb.SearchAccommodationResponse{Accommodations: []*pb.Accommodation{}}
	for _, Accommodation := range accommodations {
		accommodationPb := mapAccommodation(Accommodation)
		response.Accommodations = append(response.Accommodations, accommodationPb)

	}
	return response, nil
}
func (handler *AccommodationHandler) FilterAccommodation(ctx context.Context, request *pb.FilterAccommodationRequest) (*pb.FilterAccommodationResponse, error) {
	amenities := request.Amenities
	accommodationsPb := request.Accommodations
	var accommodationsToFilter []*domain.Accommodation
	for _, Accommodation := range accommodationsPb {
		accommodationDomain := mapAccommodationPb(Accommodation)
		accommodationsToFilter = append(accommodationsToFilter, accommodationDomain)

	}
	accommodations, err := handler.service.CheckAccommodationForAmenities(amenities, accommodationsToFilter)
	if err != nil {
		return nil, err
	}
	response := &pb.FilterAccommodationResponse{Accommodations: []*pb.Accommodation{}}
	for _, Accommodation := range accommodations {
		accommodationPb := mapAccommodation(Accommodation)
		response.Accommodations = append(response.Accommodations, accommodationPb)

	}
	return response, nil
}
func (handler *AccommodationHandler) ChangeAccommodationReservationType(ctx context.Context, request *pb.ChangeReservationTypeRequest) (*pb.ChangeReservationTypeResponse, error) {
	accommodationId := request.Id
	objectId, err := primitive.ObjectIDFromHex(accommodationId)
	confirmationType := request.ConfirmationReservationType

	if err != nil {
		return nil, err
	}
	accommodation := domain.Accommodation{
		Id:                      objectId,
		ReservationConfirmation: confirmationType,
	}
	err = handler.service.UpdateReservationConfirmationType(&accommodation)

	response := &pb.ChangeReservationTypeResponse{
		Id:  objectId.Hex(),
		Err: "nema greske",
	}
	return response, nil
}
func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	//var accommodationDTO domain.AccommodationDTO
	//fmt.Print("request: ")
	//fmt.Println(request)
	//
	//jsonBytes, err := protojson.Marshal(request)
	//
	//err = json.Unmarshal(jsonBytes, &accommodationDTO)
	//if err != nil {
	//	// Handle error
	//}
	//fmt.Println("kako se ispisati  resefvationdto")
	//fmt.Println(accommodationDTO)

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
	//layout := "2006-01-02T15:04:05.000Z"

	minGuest, err := strconv.Atoi(request.MinGuest)
	maxGuest, err := strconv.Atoi(request.MaxGuest)

	/*address := domain.Address{
		Street: accommodationDTO.Street,
		StreetNumber: accommodationDTO.StreetNumber,
		City: accommodationDTO.City,
		Country: accommodationDTO.Country,
	}*/
	objectId, err := primitive.ObjectIDFromHex(request.HostId)
	//accommodationId := uuid.New()
	var array []string
	if err != nil {
		return nil, err
	}
	//files := request.Files

	//fmt.Println(files)

	/*for _, file := range files {
		fmt.Println(os.Getwd())
		fileCreated, err := os.Create("example.txt")
		if err != nil {
			panic(err)

		}
		fmt.Println(fileCreated.Name() + "hahahahha")
		fmt.Println(fileCreated.Readdirnames(2))
		fmt.Println("dosao dovdee")

		defer fileCreated.Close()
		filePath := "C://Users/Cvetana/Desktop/cvetaVerzijamikrosevisi/cvetaVerzijamikrosevisi/microservices_demo-master/microservices_demo-master/uploads" + accommodationId.String() + file.Filename + ".jpg" // Change the extension based on the file type
		err = os.WriteFile(filePath, file.Content, 0644)
		if err != nil {
			panic(err)
		}
		array = append(array, file.Filename)
	}*/
	createAccommodation := domain.Accommodation{
		Name:                    request.Name,
		ReservationConfirmation: request.ReservationConfirmation,
		Street:                  request.Street,
		StreetNumber:            request.StreetNumber,
		City:                    request.City,
		Country:                 request.Country,
		MinGuest:                minGuest,
		MaxGuest:                maxGuest,
		HostId:                  objectId,
		Wifi:                    request.Wifi,
		Kitchen:                 request.Kitchen,
		AirConditioning:         request.AirConditioning,
		FreeParking:             request.FreeParking,
		PriceType:               request.PriceType,
		Photos:                  array,
	}

	err = handler.service.Create(&createAccommodation)
	AccommodationPb := mapAccommodation(&createAccommodation)
	response := &pb.CreateAccommodationResponse{
		Id: AccommodationPb.Id,
	}
	return response, nil
}
func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Accommodation, err := handler.service.Get(objectId)
	fmt.Println(Accommodation)
	if err != nil {
		return nil, err
	}
	AccommodationPb := mapAccommodation(Accommodation)
	fmt.Println(AccommodationPb)
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
func (handler *AccommodationHandler) GetAllByHostId(ctx context.Context, request *pb.GetByHostIdRequest) (*pb.GetAllResponse, error) {
	hostId := request.HostId
	objectId, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return nil, err
	}
	Accommodations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Accommodations: []*pb.Accommodation{},
	}
	for _, Accommodation := range Accommodations {
		if Accommodation.HostId == objectId {
			current := mapAccommodation(Accommodation)
			response.Accommodations = append(response.Accommodations, current)
		}
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
		accommodationIds = append(accommodationIds, accommodation.Id.Hex())
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
