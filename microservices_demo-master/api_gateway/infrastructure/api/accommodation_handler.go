package api

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	accommodations "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	users "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"golang.org/x/exp/slices"
	"net/http"
)

type AccommodationHandler struct {
	reservationClientAddress string
	userClientAddress        string
	accommodationAddress     string
	termAddress              string
}

func (handler *AccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/accommodations-price-range", handler.GetAccommodationsByPriceRange)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/accommodations-prominent-host", handler.GetAccommodationsByProminentHost)
	if err != nil {
		panic(err)
	}
}

func NewAccommodationHandler(reservationClientAddress, userClientAddress, accommodationAddress, termAddress string) Handler {
	return &ReservationHandler{
		reservationClientAddress: reservationClientAddress,
		userClientAddress:        userClientAddress,
		accommodationAddress:     accommodationAddress,
		termAddress:              termAddress,
	}
}

func (handler *AccommodationHandler) GetAccommodationsByPriceRange(http.ResponseWriter, *http.Request, map[string]string) {
}

func (handler *AccommodationHandler) GetAccommodationsByProminentHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userClient := services.NewUserClient(handler.userClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.userClientAddress)
	allAccommodations, err := accommodationClient.GetAll(context.TODO(), &accommodations.GetAllRequest{})
	prominentHosts, err := userClient.GetProminentHosts(context.TODO(), &users.GetProminentHostRequest{})
	var accommodationsList = make([]domain.Accommodation, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	for _, Accommodation := range allAccommodations.Accommodations {
		contains := slices.Contains(prominentHosts.HostsID, Accommodation.HostId)
		if contains {
			accommodationToAdd := domain.Accommodation{Name: Accommodation.Name, Street: Accommodation.Street}
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

func (handler *AccommodationHandler) GetAccommodationsByAmenities() {}
