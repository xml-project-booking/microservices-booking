package api

import (
	"context"
	"encoding/json"
	"errors"
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
func (handler *TermHandler) GetTermInfoByAccommodationId(ctx context.Context, request *pb.TermInfoRequest) (*pb.TermInfoResponse, error) {
	id := request.AccommodationId
	objectId, _ := primitive.ObjectIDFromHex(id)
	term, _ := handler.service.GetOneByAccommodationId(objectId)
	fmt.Println(term)
	response := &pb.TermInfoResponse{
		Price:     int64(term.Value),
		Type:      term.PriceType,
		FullPrice: 0,
	}
	fmt.Println(response)
	return response, nil
}
func (handler *TermHandler) GetAllAccommodationIdsInTimePeriod(ctx context.Context, request *pb.TimePeriodRequest) (*pb.TimePeriodResponse, error) {
	layout := "2006-01-02T15:04:05.000Z"
	startdate, _ := time.Parse(layout, request.StartDate)
	enddate, _ := time.Parse(layout, request.EndDate)
	var terms = handler.service.GetAccommodationsInTimePeriod(startdate, enddate)
	response := &pb.TimePeriodResponse{}
	for _, term := range terms {
		response.AccommodationsIds = append(response.AccommodationsIds, term.AccommodationID.Hex())
	}

	/*seen := make(map[string]bool)
	result := []string{}

	for _, val := range response.AccommodationIds {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}*/

	return response, nil
}
func (handler *TermHandler) GetAllAccommodationIdsInPriceRange(ctx context.Context, request *pb.PriceRangeRequest) (*pb.PriceRangeResponse, error) {
	minPrice := request.MinPrice
	maxPrice := request.MaxPrice
	var terms = handler.service.GetAccommodationsInPriceRange(minPrice, maxPrice)
	response := &pb.PriceRangeResponse{}
	for _, term := range terms {
		response.AccommodationIds = append(response.AccommodationIds, term.AccommodationID.Hex())
	}

	/*seen := make(map[string]bool)
	result := []string{}

	for _, val := range response.AccommodationIds {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}*/

	return response, nil
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
	if !requestIsValid(request.PriceType, int32(request.Value), accId) {
		return nil, errors.New("Cannot accomodation, value or type can't be zero")
	}

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

func (handler *TermHandler) GetByAccommodationId(ctx context.Context, request *pb.GetByAccommodationIdRequest) (*pb.GetByAccommodationIdResponse, error) {
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
	if !requestIsValid(request.PriceType, int32(request.Value), accId) {
		return nil, errors.New("Accomodation, value or type can't be zero")
	}
	usrId, _ := primitive.ObjectIDFromHex(request.UserId)

	strtDate, _ := time.Parse(layout, request.StartDate)
	eDate, _ := time.Parse(layout, request.EndDate)
	strtDate = strtDate.Truncate(24 * time.Hour)
	eDate = eDate.Truncate(24 * time.Hour)

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

func (handler *TermHandler) UpdateInPeriod(ctx context.Context, request *pb.UpdateInPeriodRequest) (*pb.UpdateResponse, error) {
	// Convert request to JSON

	jsonBytes, err := protojson.Marshal(request)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "Term Handler",
			"action":    "CRADA731",
			"timestamp": time.Now().String(),
		}).Error("Failed to cast JSON to Term!")
		return nil, err
	}

	// Unmarshal JSON to request object
	updateRequest := &pb.UpdateInPeriodRequest{}
	err = protojson.Unmarshal(jsonBytes, updateRequest)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	println("ACCID: " + updateRequest.AccommodationId)
	// Convert string IDs to primitive.ObjectIDs
	accId, _ := primitive.ObjectIDFromHex(updateRequest.AccommodationId)

	println(updateRequest.PriceType, int32(updateRequest.Value), accId.String())
	if !requestIsValid(updateRequest.PriceType, int32(updateRequest.Value), accId) {
		return nil, errors.New("Cannot accomodation, value or type can't be zero")
	}
	usrId, _ := primitive.ObjectIDFromHex(updateRequest.UserId)

	layout := "2006-01-02T15:04:05.000Z"
	// Parse start and end date strings to time.Time
	startDate, _ := time.Parse(layout, updateRequest.StartDate)
	endDate, _ := time.Parse(layout, updateRequest.EndDate)
	// Set date hours, secons ... to 0
	startDate = startDate.Truncate(24 * time.Hour)
	endDate = endDate.Truncate(24 * time.Hour)
	endDate = endDate.AddDate(0, 0, 1)
	startDate = startDate.AddDate(0, 0, 1)

	// Get all terms
	terms, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}

	// Check if any existing term in the specified period has a non-nil or non-zero user ID
	existingTerm := findTermInPeriodWithUserId(terms, accId, startDate, endDate)
	if existingTerm != nil {
		return nil, errors.New("Cannot update/create terms. An existing term in the period already has a non-nil or non-zero user ID")
	}

	// Iterate through dates in the period
	for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
		// Check if a term exists for the date and accommodation ID
		existingTerm := findTermByDateAndAccommodation(terms, date, accId)
		if existingTerm != nil {
			// Update the existing term with new information
			existingTerm.UserID = usrId
			existingTerm.PriceType = updateRequest.PriceType
			existingTerm.Value = int32(updateRequest.Value)
			err = handler.service.Update(existingTerm)
			if err != nil {
				return nil, err
			}
		} else {
			// Create a new term for the date and accommodation ID
			newTerm := domain.NewTerm(accId, usrId, updateRequest.PriceType, int32(updateRequest.Value), date)
			err = handler.service.Create(newTerm)
			if err != nil {
				return nil, err
			}
		}
	}

	// Prepare the response
	res := pb.UpdateResponse{}

	return &res, nil
}

// Helper function to check if request is valid
func findTermByDateAndAccommodation(terms []*domain.Term, date time.Time, accId primitive.ObjectID) *domain.Term {
	for _, term := range terms {
		if term.Date.Equal(date) && term.AccommodationID == accId {
			return term
		}
	}
	return nil
}

// Helper function to checkIfRequestIsValid
func requestIsValid(priceType string, value int32, accId primitive.ObjectID) bool {
	if accId.IsZero() {
		return false
	}

	if priceType == "" {
		return false
	}

	if value <= 0 {
		return false
	}

	return true
}

// Helper function to check if any existing term in the specified period has a non-nil or non-zero user ID
func findTermInPeriodWithUserId(terms []*domain.Term, accId primitive.ObjectID, startDate, endDate time.Time) *domain.Term {
	for _, term := range terms {
		if term.AccommodationID == accId && term.Date.After(startDate.AddDate(0, 0, -1)) && term.Date.Before(endDate.AddDate(0, 0, 1)) {
			if !term.UserID.IsZero() {
				return term
			}
		}
	}
	return nil
}

func (handler *TermHandler) DeleteInPeriod(ctx context.Context, request *pb.DeleteInPeriodRequest) (*pb.DeleteResponse, error) {
	// Convert request to JSON
	jsonBytes, err := protojson.Marshal(request)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "Term Handler",
			"action":    "CRADA731",
			"timestamp": time.Now().String(),
		}).Error("Failed to cast JSON to Term!")
		return nil, err
	}

	// Unmarshal JSON to request object
	deleteRequest := &pb.DeleteInPeriodRequest{}
	err = protojson.Unmarshal(jsonBytes, deleteRequest)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Convert string IDs to primitive.ObjectIDs
	accId, _ := primitive.ObjectIDFromHex(deleteRequest.AccommodationId)
	layout := "2006-01-02T15:04:05.000Z"

	// Parse start and end date strings to time.Time
	startDate, _ := time.Parse(layout, deleteRequest.StartDate)
	endDate, _ := time.Parse(layout, deleteRequest.EndDate)
	startDate = startDate.Truncate(24 * time.Hour)
	endDate = endDate.Truncate(24 * time.Hour)

	// Get all terms
	terms, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}

	// Iterate through terms, delete matching ones
	for _, term := range terms {

		if term.AccommodationID == accId && term.Date.After(startDate) && term.Date.Before(endDate) {
			err = handler.service.Delete(term)
			if err != nil {
				return nil, err
			}
		}
	}

	// Prepare the response
	res := pb.DeleteResponse{}

	return &res, nil
}

// GetAvailableAccommodationsInPeriodRequest
func (handler *TermHandler) GetAvailableAccommodationsInPeriod(ctx context.Context, request *pb.GetAvailableAccommodationsInPeriodRequest) (*pb.GetAvailableAccommodationsInPeriodResponse, error) {

	accommodationIds, err := handler.service.GetAvailableAccommodationsInPeriod(request.StartDate, request.EndDate)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	res := pb.GetAvailableAccommodationsInPeriodResponse{
		AccommodationWithPrice: accommodationIds,
	}

	return &res, nil
}
