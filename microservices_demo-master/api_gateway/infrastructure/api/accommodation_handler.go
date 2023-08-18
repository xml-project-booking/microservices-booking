package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	accommodations "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	reservations "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	terms "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	users "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"golang.org/x/exp/slices"
	"net/http"
	"strconv"
	"strings"
)

type AccommodationHandler struct {
	reservationClientAddress string
	userClientAddress        string
	accommodationAddress     string
	termAddress              string
}

func (handler *AccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/accommodations-price-range/{min}/{max}", handler.GetAccommodationsByPriceRange)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/accommodations-prominent-host", handler.GetAccommodationsByProminentHost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/search-accommodations", handler.SearchAccommodations)
	if err != nil {
		panic(err)
	}
}

func NewAccommodationHandler(reservationClientAddress, userClientAddress, accommodationAddress, termAddress string) Handler {
	return &AccommodationHandler{
		reservationClientAddress: reservationClientAddress,
		userClientAddress:        userClientAddress,
		accommodationAddress:     accommodationAddress,
		termAddress:              termAddress,
	}
}

func (handler *AccommodationHandler) GetAccommodationsByPriceRange(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	termClient := services.NewTermClient(handler.termAddress)
	minPrice, _ := strconv.Atoi(pathParams["min"])
	maxPrice, _ := strconv.Atoi(pathParams["max"])
	decoder := json.NewDecoder(r.Body)
	var t []domain.Accommodation
	err := decoder.Decode(&t)
	fmt.Println(t)
	if err != nil {
		panic(err)

	}
	accommodationsIds, err := termClient.GetAllAccommodationIdsInPriceRange(context.TODO(), &terms.PriceRangeRequest{MinPrice: int32(minPrice), MaxPrice: int32(maxPrice)})
	fmt.Println(accommodationsIds)
	var accommodationsList = make([]domain.Accommodation, 0)
	for _, Accommodation := range t {
		contains := slices.Contains(accommodationsIds.AccommodationIds, Accommodation.Id.Hex())
		if contains {
			accommodationToAdd := domain.Accommodation{Name: Accommodation.Name, Street: Accommodation.Street, ReservationConfirmation: Accommodation.ReservationConfirmation, City: Accommodation.City,
				StreetNumber: Accommodation.StreetNumber, Wifi: Accommodation.Wifi, Kitchen: Accommodation.Kitchen, AirConditioning: Accommodation.AirConditioning, MinGuest: Accommodation.MinGuest,
				MaxGuest: Accommodation.MaxGuest, FreeParking: Accommodation.FreeParking, Country: Accommodation.Country, Id: Accommodation.Id}
			accommodationsList = append(accommodationsList, accommodationToAdd)
		}

	}
	response, err := json.Marshal(accommodationsList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *AccommodationHandler) GetAccommodationsByProminentHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userClient := services.NewUserClient(handler.userClientAddress)
	decoder := json.NewDecoder(r.Body)
	var t []domain.Accommodation
	err := decoder.Decode(&t)
	fmt.Println(t)
	if err != nil {
		panic(err)

	}
	prominentHosts, err := userClient.GetProminentHosts(context.TODO(), &users.GetProminentHostRequest{})
	var accommodationsList = make([]domain.Accommodation, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	for _, Accommodation := range t {
		contains := slices.Contains(prominentHosts.HostsID, Accommodation.HostId.Hex())
		if contains {
			accommodationToAdd := domain.Accommodation{Name: Accommodation.Name, Street: Accommodation.Street, ReservationConfirmation: Accommodation.ReservationConfirmation, City: Accommodation.City,
				StreetNumber: Accommodation.StreetNumber, Wifi: Accommodation.Wifi, Kitchen: Accommodation.Kitchen, AirConditioning: Accommodation.AirConditioning, MinGuest: Accommodation.MinGuest,
				MaxGuest: Accommodation.MaxGuest, FreeParking: Accommodation.FreeParking, Country: Accommodation.Country, Id: Accommodation.Id}
			accommodationsList = append(accommodationsList, accommodationToAdd)
		}

	}

	response, err := json.Marshal(accommodationsList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *AccommodationHandler) SearchAccommodations(w http.ResponseWriter, request *http.Request, params map[string]string) {
	accommodationClient := services.NewAccommodationClient(handler.accommodationAddress)
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	decoder := json.NewDecoder(request.Body)
	var t domain.SearchDTO
	err := decoder.Decode(&t)
	fmt.Println(t)
	if err != nil {
		panic(err)

	}
	response, err := reservationClient.GetAccommodationsReservedInTimePeriod(context.TODO(), &reservations.GetAccTimePeriodRequest{
		StartDate: t.StartDate, EndDate: t.EndDate,
	})
	if err != nil {
		panic(err)
	}
	allAccommodations, err := accommodationClient.GetAll(context.TODO(), &accommodations.GetAllRequest{})
	searchAccommodations := make([]domain.Accommodation, 0)
	for _, Accommodation := range allAccommodations.Accommodations {
		contains := slices.Contains(response.AccommodationIds, Accommodation.Id)
		if !contains {
			accommodationToAdd := domain.Accommodation{Name: Accommodation.Name, Street: Accommodation.Street, City: Accommodation.City,
				Wifi: Accommodation.Wifi, Kitchen: Accommodation.Kitchen, AirConditioning: Accommodation.AirConditioning, MinGuest: int(Accommodation.MinGuest),
				MaxGuest: int(Accommodation.MaxGuest), FreeParking: Accommodation.FreeParking, Country: Accommodation.Country}
			searchAccommodations = append(searchAccommodations, accommodationToAdd)
		}
	}

}
func (handler *AccommodationHandler) SearchAccommodationsByLocationAndGuestNumber(accommodationsList []domain.Accommodation, location string, guestNumber int) []domain.Accommodation {
	var filteredList []domain.Accommodation
	for _, Accommodation := range accommodationsList {
		result := strings.Contains(strings.ToLower(Accommodation.Country), strings.ToLower(location))
		resultOne := strings.Contains(strings.ToLower(Accommodation.City), strings.ToLower(location))
		if (result || resultOne) && (guestNumber >= Accommodation.MinGuest && guestNumber <= Accommodation.MaxGuest) {
			filteredList = append(filteredList, Accommodation)
		}
	}
	return filteredList

}
