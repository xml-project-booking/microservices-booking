package api

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	accommodationGw "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	reservationGw "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	userGw "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	//"log"
	"net/http"
)

type UserHandler struct {
	UserAddress          string
	ReservationAddress   string
	AccommodationAddress string
}

func NewUserHandler(userAddress string, reservationAddress string, accommodationAddress string) Handler {
	return &UserHandler{
		UserAddress:          userAddress,
		ReservationAddress:   reservationAddress,
		AccommodationAddress: accommodationAddress,
	}
}
func (handler *UserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/users/deleted/{id}", handler.DeleteUser)
	if err != nil {
		panic(err)
	}
}
func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userDetails := &domain.UserInfo{Id: id}
	user, err := handler.getUserById(userDetails)
	if user.Role == "GUEST" {
		hasReservation, err := handler.hasReservationsGuest(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if hasReservation {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	} else if user.Role == "HOST" {
		hasReservation, err := handler.hasReservationsHost(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if hasReservation {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = handler.deleteAccomodations(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	err = handler.deleteAccount(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (handler *UserHandler) getUserById(userDetails *domain.UserInfo) (userDetailsOutput *domain.UserInfo, err error) {
	userClient := services.NewUserClient(handler.UserAddress)
	user, err := userClient.Get(context.TODO(), &userGw.GetRequest{Id: userDetails.Id})
	if err != nil {
		return nil, err
	}
	stringRole := "HOST"
	if user.User.IsHost == false {
		stringRole = "GUEST"
	}
	userDetailsOutput = &domain.UserInfo{Id: userDetails.Id}
	userDetailsOutput.Role = stringRole
	if err != nil {
		return nil, err
	}
	return userDetailsOutput, nil
}
func (handler *UserHandler) deleteAccount(id string) error {
	userClient := services.NewUserClient(handler.UserAddress)
	deleteResponse, err := userClient.Delete(context.TODO(), &userGw.DeleteRequest{
		Id: id,
	})
	if err != nil {
		return err
	}
	if deleteResponse.RequestResult.Code != 200 {
		return errors.New(deleteResponse.RequestResult.Message)
	}
	return nil
}
func (handler *UserHandler) hasReservationsGuest(id string) (bool, error) {
	reservationClient := services.NewReservationClient(handler.ReservationAddress)
	reservationsResponse, err := reservationClient.HasActiveReservations(context.TODO(), &reservationGw.HasActiveReservationsRequest{
		Id: id,
	})
	if err != nil {
		return false, err
	}
	return reservationsResponse.HasReservations, nil
}
func (handler *UserHandler) hasReservationsHost(id string) (bool, error) {
	accomodationClient := services.NewAccommodationClient(handler.AccommodationAddress)
	accommodationIds, err := accomodationClient.GetAllIdsByHost(context.TODO(), &accommodationGw.GetAllIdsByHostRequest{
		Id: id,
	})
	if err != nil {
		return false, err
	}
	reservationClient := services.NewReservationClient(handler.ReservationAddress)
	reservationsResponse, err2 := reservationClient.GetAllFuture(context.TODO(), &reservationGw.GetAllFutureRequest{})
	if err2 != nil {
		return false, err2
	}
	hasReservations := false
	for _, reservation := range reservationsResponse.Reservations {
		for _, idA := range accommodationIds.Ids {
			if reservation.AccommodationID == idA && reservation.IsConfirmed == true {
				hasReservations = true
			}
		}
	}
	return hasReservations, nil
}
func (handler *UserHandler) deleteAccomodations(id string) error {
	accommodationClient := services.NewAccommodationClient(handler.AccommodationAddress)
	deleteResponse, err := accommodationClient.DeleteAllByHost(context.TODO(), &accommodationGw.DeleteAllByHostRequest{
		Id: id,
	})
	if err != nil {
		return err
	}
	if deleteResponse.RequestResult.Code != 200 {
		return errors.New(deleteResponse.RequestResult.Message)
	}
	return nil
}
