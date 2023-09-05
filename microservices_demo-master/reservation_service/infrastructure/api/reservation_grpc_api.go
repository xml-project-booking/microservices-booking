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

func (handler ReservationHandler) GetAccommodationsReservedInTimePeriod(ctx context.Context, request *pb.GetAccTimePeriodRequest) (*pb.GetAccTimePeriodResponse, error) {
	layout := "2006-01-02T15:04:05.000Z"
	startDate, _ := time.Parse(layout, request.StartDate)
	endDate, _ := time.Parse(layout, request.EndDate)
	reservations := handler.service.GetAllReservationsSearch(startDate, endDate)
	var accommodationsIds []string
	for _, Reservation := range reservations {
		accommodationsIds = append(accommodationsIds, Reservation.AccommodationID.Hex())
	}
	response := &pb.GetAccTimePeriodResponse{AccommodationIds: accommodationsIds}
	return response, nil
}
func (handler *ReservationHandler) CheckReservationRequirementsHost(ctx context.Context, request *pb.ReservationRequirementsHostRequest) (*pb.ReservationRequirementsHostResponse, error) {
	hostId := request.HostId
	objectId, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return nil, err
	}
	a := handler.service.CheckReservationNumberForHost(objectId)
	b := handler.service.CheckCancellationRate(objectId)
	c := handler.service.CheckTotalReservationDuration(objectId)
	if a && b && c {
		response := &pb.ReservationRequirementsHostResponse{IsPossible: true}
		return response, nil
	} else {
		response := &pb.ReservationRequirementsHostResponse{IsPossible: false}
		return response, nil
	}
}
func (handler *ReservationHandler) CancelReservation(ctx context.Context, request *pb.CancelReservationRequest) (*pb.CancelReservationResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	Reservation, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	isDeletedRes := handler.service.CancelReservation(Reservation)
	if err != nil {
		return nil, err
	}
	//isDeletedPb := mapReservationCancelation(isDeletedRes)
	response := &pb.CancelReservationResponse{
		IsDeleted: isDeletedRes,
	}
	return response, nil

}

func (handler *ReservationHandler) TermCheck(ctx context.Context, request *pb.TermCheckRequest) (*pb.TermCheckResponse, error) {
	var termCheckDTO domain.TermCheckDTO
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

	err = json.Unmarshal(jsonBytes, &termCheckDTO)
	if err != nil {
		// Handle error
	}

	layout := "2006-01-02T15:04:05.000Z"
	startDate, _ := time.Parse(layout, termCheckDTO.StartDate)
	endDate, _ := time.Parse(layout, termCheckDTO.EndDate)
	accommodationId, err := primitive.ObjectIDFromHex(termCheckDTO.Id)
	fmt.Println("ovo se dobije za acc id")
	fmt.Println(accommodationId)
	getId := handler.service.CheckForReservationInDateRangeAndGetUserId(accommodationId, startDate, endDate)
	fmt.Println(getId + "ovooo sam dobilaaaa")
	res := &pb.TermCheckResponse{
		HasReservation: getId,
	}
	fmt.Println()
	fmt.Println(res)
	return res, nil
}

func (handler *ReservationHandler) DeleteReservationRequestGuest(ctx context.Context, request *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	boolVar := handler.service.Delete(objectId)
	if err != nil {
		return nil, err
	}
	isDeleted := strconv.FormatBool(boolVar)
	//ReservationPb := mapReservation(Reservation)
	response := &pb.DeleteReservationResponse{
		Id: isDeleted,
	}
	return response, nil
}

func (handler *ReservationHandler) MakeRequestForReservation(ctx context.Context, request *pb.ReservationRequest) (*pb.ReservationRequestResponse, error) {

	var reservationDTO domain.ReservationDTO

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
	layout := "2006-01-02T15:04:05.000Z"
	startDate, _ := time.Parse(layout, reservationDTO.StartDate)
	endDate, _ := time.Parse(layout, reservationDTO.EndDate)
	num, err := strconv.Atoi(reservationDTO.GuestNumber)
	Minnum, err := strconv.Atoi(reservationDTO.MinGuest)
	MaxNum, err := strconv.Atoi(reservationDTO.MaxGuest)
	b, err := strconv.ParseBool("false")
	var isValidGuestNum = handler.service.CheckIfNumberOfGuestIsValid(int64(Minnum), int64(MaxNum), int64(num))
	re := &pb.ReservationRequestResponse{Id: "nedozvoljen broj gostiju"}
	fmt.Print("ovoooo je vrednost")
	fmt.Println(isValidGuestNum)
	if !(isValidGuestNum) {
		return re, nil
	}
	var isTaken = handler.service.CheckForReservationInDateRange(reservationDTO.AccommodationID, startDate, endDate)
	re = &pb.ReservationRequestResponse{Id: "greska"}
	if isTaken == true {
		return re, nil
	}

	reservationRequest := domain.Reservation{
		AccommodationID:   reservationDTO.AccommodationID,
		StartDate:         startDate,
		EndDate:           endDate,
		Confirmation:      b,
		GuestNumber:       int64(num),
		GuestId:           reservationDTO.GuestId,
		ReservationStatus: "PENDING",
		HostId:            reservationDTO.HostId,
	}

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
func (handler *ReservationHandler) GetAllByAccommodation(ctx context.Context, request *pb.GetAllByAccommodationRequest) (*pb.GetAllByAccommodationResponse, error) {
	accommodationId := request.Id
	fmt.Println("ov je id od accommodationa")
	fmt.Println(accommodationId)
	objectId, err := primitive.ObjectIDFromHex(accommodationId)
	Reservations, err := handler.service.GetAllReservationsByAccommodation(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllByAccommodationResponse{
		Reservations: []*pb.Reservation{},
	}
	fmt.Println("ovo su rezevracije")

	for _, Reservation := range Reservations {
		fmt.Println(Reservation)
		current := mapReservation(Reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetAllByAccommodationConfirmed(ctx context.Context, request *pb.GetAllByAccommodationRequest) (*pb.GetAllByAccommodationResponse, error) {
	accommodationId := request.Id

	objectId, err := primitive.ObjectIDFromHex(accommodationId)
	Reservations, err := handler.service.GetAllConfirmedReservationsByAccommodation(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllByAccommodationResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, Reservation := range Reservations {
		current := mapReservation(Reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetAllByGuest(ctx context.Context, request *pb.GetAllByGuestRequest) (*pb.GetAllByAccommodationResponse, error) {
	guestId := request.Id
	fmt.Println(guestId)
	objectId, err := primitive.ObjectIDFromHex(guestId)
	fmt.Println("ajdeee")
	Reservations, err := handler.service.GetAllReservationsByGuestId(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllByAccommodationResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, Reservation := range Reservations {
		current := mapReservation(Reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetAllByGuestPending(ctx context.Context, request *pb.GetAllByGuestRequest) (*pb.GetAllByAccommodationResponse, error) {
	guestId := request.Id
	fmt.Println(guestId)
	objectId, err := primitive.ObjectIDFromHex(guestId)
	fmt.Println("ajdeee")
	Reservations, err := handler.service.GetAllReservationsByGuestIdPending(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllByAccommodationResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, Reservation := range Reservations {
		current := mapReservation(Reservation)
		response.Reservations = append(response.Reservations, current)
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
func (handler *ReservationHandler) GetAllFuture(ctx context.Context, request *pb.GetAllFutureRequest) (*pb.GetAllFutureResponse, error) {
	reservations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}

	allReservations := []*domain.Reservation{}
	for _, reservation := range reservations {
		if reservation.StartDate.After(time.Now()) {
			allReservations = append(allReservations, reservation)
		}
	}
	response := &pb.GetAllFutureResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, r := range allReservations {
		current := mapReservation(r)
		response.Reservations = append(response.Reservations, current)
	}

	return response, nil
}
func (handler *ReservationHandler) HasActiveReservations(ctx context.Context, request *pb.HasActiveReservationsRequest) (response *pb.HasActiveReservationsResponse, err error) {
	id := request.Id
	guestId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	reservations, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}

	//ako se promeni na enum ovde menjaj
	allReservations := []*domain.Reservation{}
	for _, reservation := range reservations {
		if reservation.GuestId == guestId && reservation.ReservationStatus == "CONFIRMED" {
			allReservations = append(allReservations, reservation)
		}
	}
	hasReservations := false
	if len(allReservations) != 0 {
		hasReservations = true
	}

	// Prepare the response
	response = &pb.HasActiveReservationsResponse{
		HasReservations: hasReservations,
	}
	return response, nil
}

func (handler *ReservationHandler) ConfirmReservationAutomatically(ctx context.Context, request *pb.ReservationRequest) (*pb.ConfirmReservationAutomaticallyMessage, error) {
	var reservationDTO domain.ReservationDTO
	fmt.Print("request: ")
	fmt.Println(request)
	//id := request.Id
	//objectId, err := primitive.ObjectIDFromHex(id)

	jsonBytes, err := protojson.Marshal(request)

	err = json.Unmarshal(jsonBytes, &reservationDTO)
	if err != nil {
		// Handle error
	}
	//Reservation, err := handler.service.Get(objectId)
	layout := "2006-01-02T15:04:05.000Z"
	startDate, _ := time.Parse(layout, reservationDTO.StartDate)
	endDate, _ := time.Parse(layout, reservationDTO.EndDate)
	num, err := strconv.Atoi(reservationDTO.GuestNumber)
	Minnum, err := strconv.Atoi(reservationDTO.MinGuest)
	MaxNum, err := strconv.Atoi(reservationDTO.MaxGuest)
	//b, err := strconv.ParseBool("true")
	var isValidGuestNum = handler.service.CheckIfNumberOfGuestIsValid(int64(Minnum), int64(MaxNum), int64(num))
	result := &pb.ConfirmReservationAutomaticallyMessage{Id: "nedozvoljen broj gostiju"}
	fmt.Println("vrednost promenjive isvalidguestnum")
	fmt.Println(isValidGuestNum)
	if !(isValidGuestNum) {
		fmt.Println("uuuuuuuuuuuuuuuuuuuuuuu")
		return result, nil
	}
	var isTaken = handler.service.CheckForReservationInDateRange(reservationDTO.AccommodationID, startDate, endDate)
	//re := &pb.ReservationRequestResponse{Id: "greska"}
	fmt.Println(isTaken)
	if isTaken == true {

		//err = handler.service.UpdateReservationStatusForCanceled(&reservationRequest)
		res := &pb.ConfirmReservationAutomaticallyMessage{
			Id: "greska",
		}
		return res, nil
	}
	reservationRequest := domain.Reservation{
		AccommodationID:   reservationDTO.AccommodationID,
		StartDate:         startDate,
		EndDate:           endDate,
		Confirmation:      true,
		GuestNumber:       int64(num),
		GuestId:           reservationDTO.GuestId,
		ReservationStatus: "CONFIRMED",
		HostId:            reservationDTO.HostId,
	}

	err = handler.service.Create(&reservationRequest)

	err = handler.service.UpdateReservationStatusForConfirmed(&reservationRequest)
	//var listOfReservationRequests = handler.service.GetAllReservationRequestWhichOverlapsWithConfirmed(Reservation.StartDate, Reservation.EndDate, Reservation.AccommodationID)
	//handler.service.CancelAllReservationRequestOverlaping(listOfReservationRequests)
	ReservationPb := mapReservation(&reservationRequest)
	response := &pb.ConfirmReservationAutomaticallyMessage{
		Id: ReservationPb.Id,
	}
	return response, nil
}
func (handler *ReservationHandler) ConfirmReservationManually(ctx context.Context, request *pb.ConfirmReservationManuallyRequest) (*pb.ConfirmReservationManuallyResponse, error) {
	var reservationDTO domain.ReservationDTO
	fmt.Print("request: ")

	fmt.Println(request)
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)

	jsonBytes, err := protojson.Marshal(request)

	err = json.Unmarshal(jsonBytes, &reservationDTO)
	if err != nil {
		// Handle error
	}

	//layout := "2006-01-02T15:04:05.000Z"
	/*startDate, _ := time.Parse(layout, reservationDTO.StartDate)
	endDate, _ := time.Parse(layout, reservationDTO.EndDate)
	num, err := strconv.Atoi(reservationDTO.GuestNumber)
	b, err := strconv.ParseBool(reservationDTO.Confirmation)*/
	Reservation, err := handler.service.Get(objectId)

	//err = handler.service.Create(&reservationRequest)
	err = handler.service.UpdateReservationStatusForConfirmed(Reservation)
	var listOfReservationRequests = handler.service.GetAllReservationRequestWhichOverlapsWithConfirmed(Reservation.StartDate, Reservation.EndDate, Reservation.AccommodationID)
	fmt.Println("vo je duzina liste sa prekalpajucim terminima")
	fmt.Println(listOfReservationRequests)
	handler.service.CancelAllReservationRequestOverlaping(listOfReservationRequests)
	ReservationPb := mapReservation(Reservation)
	response := &pb.ConfirmReservationManuallyResponse{
		Id: ReservationPb.Id,
	}
	return response, nil
}
func (handler *ReservationHandler) CancelReservationManually(ctx context.Context, request *pb.CancelReservationManuallyRequest) (*pb.CancelReservationManuallyResponse, error) {
	var reservationDTO domain.ReservationDTO
	fmt.Print("request: ")
	fmt.Println(request)
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	jsonBytes, err := protojson.Marshal(request)

	err = json.Unmarshal(jsonBytes, &reservationDTO)
	if err != nil {
		// Handle error
	}

	/*layout := "2006-01-02T15:04:05.000Z"
	startDate, _ := time.Parse(layout, reservationDTO.StartDate)
	endDate, _ := time.Parse(layout, reservationDTO.EndDate)
	num, err := strconv.Atoi(reservationDTO.GuestNumber)
	b, err := strconv.ParseBool(reservationDTO.Confirmation)*/
	Reservation, err := handler.service.Get(objectId)

	/*reservationRequest := domain.Reservation{
		AccommodationID:   reservationDTO.AccommodationID,
		StartDate:         startDate,
		EndDate:           endDate,
		Confirmation:      b,
		GuestNumber:       int64(num),
		GuestId:           reservationDTO.GuestId,
		ReservationStatus: "CONFIRMED",
	}*/
	err = handler.service.UpdateReservationStatusForCanceled(Reservation)
	if err != nil {
		return nil, err
	}
	//err = handler.service.Create(&reservationRequest)
	//var listOfReservationRequests = handler.service.GetAllReservationRequestWhichOverlapsWithConfirmed(Reservation.StartDate, Reservation.EndDate, Reservation.AccommodationID)
	//handler.service.CancelAllReservationRequestOverlaping(listOfReservationRequests)
	ReservationPb := mapReservation(Reservation)
	response := &pb.CancelReservationManuallyResponse{
		Id: ReservationPb.Id,
	}
	return response, nil
}
