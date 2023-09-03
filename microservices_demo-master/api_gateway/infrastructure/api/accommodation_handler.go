package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	accommodations "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	terms "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	users "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
	"time"
)

type AccommodationHandler struct {
	reservationClientAddress string
	userClientAddress        string
	accommodationAddress     string
	termAddress              string
}

func (handler *AccommodationHandler) Init(mux *runtime.ServeMux) {
	//err := mux.HandlePath("POST", "/accommodations-price-range/{min}/{max}", handler.GetAccommodationsByPriceRange)
	//if err != nil {
	//panic(err)
	//}
	/*err = mux.HandlePath("POST", "/accommodations-prominent-host", handler.GetAccommodationsByProminentHost)
	if err != nil {
		panic(err)
	}*/
	err := mux.HandlePath("POST", "/search-accommodations", handler.SearchAccommodations)
	if err != nil {
		panic(err)
	}
	//err = mux.HandlePath("POST", "/filter-accommodations", handler.GetAccommodationsByAmenities)
	//if err != nil {
	//panic(err)
	//}
	err = mux.HandlePath("POST", "/filter-accommodations", handler.FilterAccommodation)
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

func (handler *AccommodationHandler) FilterAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	decoder := json.NewDecoder(r.Body)
	var t domain.FilterParameters

	err := decoder.Decode(&t)
	fmt.Println(t)
	if err != nil {
		panic(err)
	}
	var commonElements []*domain.Accommodation
	var commonElementsTwo []*domain.Accommodation
	for _, accommodation := range t.Accommodations {
		if accommodation.Wifi == t.Amenities[0] && accommodation.Kitchen == t.Amenities[1] &&
			accommodation.AirConditioning == t.Amenities[2] && accommodation.FreeParking == t.Amenities[3] &&
			accommodation.Price >= t.MinPrice && accommodation.Price <= t.MaxPrice {
			commonElements = append(commonElements, accommodation)
		}

	}
	if t.IsHost {
		accommodationsThird := handler.GetAccommodationsByProminentHost(t.Accommodations)
		for _, a := range commonElements {
			for _, b := range accommodationsThird {
				if a.HostId.Hex() == b {
					commonElementsTwo = append(commonElementsTwo, a)
				}
			}
		}
		response, err := json.Marshal(commonElementsTwo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	} else {

		response, err := json.Marshal(commonElements)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}
func (handler *AccommodationHandler) GetAccommodationsByAmenities(accommodationList []*domain.Accommodation, amenities []bool) []domain.Accommodation {
	termClient := services.NewTermClient(handler.termAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationAddress)
	var listPB []*accommodations.Accommodation
	for _, Acc := range accommodationList {
		acc := mapAccommodation(Acc)
		listPB = append(listPB, acc)
	}
	accomm, err := accommodationClient.FilterAccommodation(context.TODO(), &accommodations.FilterAccommodationRequest{Amenities: amenities,
		Accommodations: listPB})
	if err != nil {
		panic(err)
	}
	var accommodationsList = make([]domain.Accommodation, 0)
	for _, Accommodation := range accomm.Accommodations {
		termInfo, err := termClient.GetTermInfoByAccommodationId(context.TODO(), &terms.TermInfoRequest{AccommodationId: Accommodation.Id})
		if err != nil {
			panic(err)
		}
		objectId, _ := primitive.ObjectIDFromHex(Accommodation.Id)
		accommodationToAdd := domain.Accommodation{Name: Accommodation.Name, Street: Accommodation.Street, City: Accommodation.City,
			Wifi: Accommodation.Wifi, Kitchen: Accommodation.Kitchen, AirConditioning: Accommodation.AirConditioning, MinGuest: int(Accommodation.MinGuest),
			MaxGuest: int(Accommodation.MaxGuest), FreeParking: Accommodation.FreeParking, Country: Accommodation.Country, Id: objectId,
			Price: termInfo.Price, TotalPrice: 0, Type: termInfo.Type}
		accommodationsList = append(accommodationsList, accommodationToAdd)

	}
	return accommodationsList

}
func (handler *AccommodationHandler) GetAccommodationsByPriceRange(accommodationsToFilter []*domain.Accommodation, minPrice, maxPrice int) []string {
	termClient := services.NewTermClient(handler.termAddress)

	//var t []domain.Accommodation
	//err := decoder.Decode(&t)
	//fmt.Println(t)
	//if err != nil {
	//panic(err)

	//}

	accommodationsIds, err := termClient.GetAllAccommodationIdsInPriceRange(context.TODO(), &terms.PriceRangeRequest{MinPrice: int32(minPrice), MaxPrice: int32(maxPrice)})
	fmt.Println(accommodationsIds)
	if err != nil {
		panic(err)
	}
	//var accommodationsList = make([]domain.Accommodation, 0)
	/*for _, Accommodation := range t {
		//contains := slices.Contains(accommodationsIds.AccommodationIds, Accommodation.Id.Hex())
		//if contains {
			accommodationToAdd := domain.Accommodation{Name: Accommodation.Name, Street: Accommodation.Street, ReservationConfirmation: Accommodation.ReservationConfirmation, City: Accommodation.City,
				StreetNumber: Accommodation.StreetNumber, Wifi: Accommodation.Wifi, Kitchen: Accommodation.Kitchen, AirConditioning: Accommodation.AirConditioning, MinGuest: Accommodation.MinGuest,
				MaxGuest: Accommodation.MaxGuest, FreeParking: Accommodation.FreeParking, Country: Accommodation.Country, Id: Accommodation.Id,
				Price: Accommodation.Price, TotalPrice: Accommodation.TotalPrice, Type: Accommodation.Type}
			accommodationsList = append(accommodationsList, accommodationToAdd)
		//}

	}*/
	return accommodationsIds.AccommodationIds

}

func (handler *AccommodationHandler) GetAccommodationsByProminentHost(accommodationsToFilter []*domain.Accommodation) []string {
	userClient := services.NewUserClient(handler.userClientAddress)

	/*decoder := json.NewDecoder(r.Body)
	var t []domain.Accommodation
	err := decoder.Decode(&t)
	fmt.Println(t)
	if err != nil {
		panic(err)

	}*/
	prominentHosts, err := userClient.GetProminentHosts(context.TODO(), &users.GetProminentHostRequest{})
	//var accommodationsList = make([]domain.Accommodation, 0)
	if err != nil {
		panic(err)
	}

	/*for _, Accommodation := range t {
		contains := slices.Contains(prominentHosts.HostsID, Accommodation.HostId.Hex())
		if contains {

			accommodationToAdd := domain.Accommodation{Name: Accommodation.Name, Street: Accommodation.Street, ReservationConfirmation: Accommodation.ReservationConfirmation, City: Accommodation.City,
				StreetNumber: Accommodation.StreetNumber, Wifi: Accommodation.Wifi, Kitchen: Accommodation.Kitchen, AirConditioning: Accommodation.AirConditioning, MinGuest: Accommodation.MinGuest,
				MaxGuest: Accommodation.MaxGuest, FreeParking: Accommodation.FreeParking, Country: Accommodation.Country, Id: Accommodation.Id,
				Price: Accommodation.Price, TotalPrice: Accommodation.TotalPrice, Type: Accommodation.Type}
			accommodationsList = append(accommodationsList, accommodationToAdd)
		}

	}*/

	/*response, err := json.Marshal(accommodationsList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)*/
	return prominentHosts.HostsID

}

func (handler *AccommodationHandler) SearchAccommodations(w http.ResponseWriter, request *http.Request, params map[string]string) {
	accommodationClient := services.NewAccommodationClient(handler.accommodationAddress)
	termClient := services.NewTermClient(handler.termAddress)
	decoder := json.NewDecoder(request.Body)
	var t domain.SearchDTO
	err := decoder.Decode(&t)
	fmt.Println(t)
	startDate, err := time.Parse("2006-01-02T15:04:05Z", t.StartDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}
	endDate, err := time.Parse("2006-01-02T15:04:05Z", t.EndDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}
	duration := endDate.Sub(startDate)
	daysDifference := int(duration.Hours() / 24)
	if err != nil {
		panic(err)

	}

	if err != nil {
		panic(err)
	}
	response, err := termClient.GetAllAccommodationIdsInTimePeriod(context.TODO(), &terms.TimePeriodRequest{StartDate: t.StartDate, EndDate: t.EndDate})
	searchAccommodations := make([]domain.Accommodation, 0)
	for _, id := range response.AccommodationsIds {
		fmt.Println(id)
		res, err := accommodationClient.Get(context.TODO(), &accommodations.GetRequest{Id: id})
		if err != nil {
			panic(err)
		}
		termInfo, err := termClient.GetTermInfoByAccommodationId(context.TODO(), &terms.TermInfoRequest{AccommodationId: id})
		if err != nil {
			panic(err)
		}
		var totalPrice = 0
		if termInfo.Type == "per person" {
			totalPrice = int(termInfo.Price) * t.GuestNumber * daysDifference
		} else {
			totalPrice = int(termInfo.Price) * daysDifference
		}
		objectId, _ := primitive.ObjectIDFromHex(res.Accommodation.Id)
		accommodationToAdd := domain.Accommodation{Name: res.Accommodation.Name, Street: res.Accommodation.Street, City: res.Accommodation.City, Id: objectId,
			Wifi: res.Accommodation.Wifi, Kitchen: res.Accommodation.Kitchen, AirConditioning: res.Accommodation.AirConditioning, MinGuest: int(res.Accommodation.MinGuest),
			MaxGuest: int(res.Accommodation.MaxGuest), FreeParking: res.Accommodation.FreeParking, Country: res.Accommodation.Country,
			Price: termInfo.Price, Type: termInfo.Type, TotalPrice: int64(totalPrice)}
		searchAccommodations = append(searchAccommodations, accommodationToAdd)

	}
	searchAccommodations = handler.SearchAccommodationsByLocationAndGuestNumber(searchAccommodations, t.Location, t.GuestNumber)

	searchAccommodations = handler.GetPriceForAccommodations(searchAccommodations, t.GuestNumber)
	fmt.Println(searchAccommodations)
	responseOne, err := json.Marshal(searchAccommodations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseOne)

}
func (handler *AccommodationHandler) SearchAccommodationsByLocationAndGuestNumber(accommodationsList []domain.Accommodation, location string, guestNumber int) []domain.Accommodation {
	var filteredList []domain.Accommodation
	for _, Accommodation := range accommodationsList {
		result := strings.Contains(strings.ToLower(Accommodation.Country), strings.ToLower(location))
		fmt.Println(result)
		resultOne := strings.Contains(strings.ToLower(Accommodation.City), strings.ToLower(location))
		if (result || resultOne) && (guestNumber >= Accommodation.MinGuest && guestNumber <= Accommodation.MaxGuest) {
			filteredList = append(filteredList, Accommodation)
		}
	}
	return filteredList

}

func (handler *AccommodationHandler) GetPriceForAccommodations(accommodations []domain.Accommodation, guestNumber int) []domain.Accommodation {
	termClient := services.NewTermClient(handler.termAddress)
	for _, Accomodation := range accommodations {
		fmt.Println(Accomodation.Id.Hex())
		termInfo, err := termClient.GetTermInfoByAccommodationId(context.TODO(), &terms.TermInfoRequest{AccommodationId: Accomodation.Id.Hex()})
		if err != nil {
			panic(err)
		}
		fmt.Println(termInfo)
		Accomodation.Price = termInfo.Price
		fmt.Println(Accomodation.Price)
		Accomodation.TotalPrice = termInfo.Price * int64(guestNumber)
		Accomodation.Type = termInfo.Type
	}
	fmt.Println(accommodations)
	return accommodations

}
