package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	accommodationGw "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	ratingGw "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	reservationGw "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	userGw "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	//"log"
	"net/http"
)

type UserHandler struct {
	UserAddress          string
	ReservationAddress   string
	AccommodationAddress string
	RatingAddress        string
}

func NewUserHandler(userAddress string, reservationAddress string, accommodationAddress, ratingAddress string) Handler {
	return &UserHandler{
		UserAddress:          userAddress,
		ReservationAddress:   reservationAddress,
		AccommodationAddress: accommodationAddress,
		RatingAddress:        ratingAddress,
	}
}
func (handler *UserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/users/deleted/{id}", handler.DeleteUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/users/prominent-host/{id}", handler.ProminentHost)
	err = mux.HandlePath("GET", "/users/ratings/host/{id}", handler.RatingsHost)
	err = mux.HandlePath("GET", "/users/ratings/accommodation/{id}", handler.RatingsAccommodation)
}
func (handler *UserHandler) RatingsHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hostId := pathParams["id"]
	var ratingsHost []domain.RatingDTO
	ratingClient := services.NewRatingClient(handler.RatingAddress)
	userClient := services.NewUserClient(handler.UserAddress)
	ratings, err := ratingClient.GetRatingsByType(context.TODO(), &ratingGw.GetRatingsByTypeRequest{
		Type: 1,
		Id:   hostId,
	})
	for _, Rating := range ratings.Ratings {
		User, err := userClient.Get(context.TODO(), &userGw.GetRequest{Id: Rating.UserID})
		if err != nil {
			panic(err)
		}

		accommodationToAdd := domain.RatingDTO{
			UserID:       Rating.UserID,
			TargetId:     Rating.TargetId,
			RatingValue:  int(Rating.RatingValue),
			LastModified: Rating.LastModified,
			Name:         User.User.Name,
			Username:     User.User.Username,
			Surname:      User.User.Surname,
		}
		ratingsHost = append(ratingsHost, accommodationToAdd)
	}

	responseOne, err := json.Marshal(ratingsHost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseOne)
}
func (handler *UserHandler) RatingsAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hostId := pathParams["id"]
	var ratingsHost []domain.RatingDTO
	ratingClient := services.NewRatingClient(handler.RatingAddress)
	userClient := services.NewUserClient(handler.UserAddress)
	ratings, err := ratingClient.GetRatingsByType(context.TODO(), &ratingGw.GetRatingsByTypeRequest{
		Type: 0,
		Id:   hostId,
	})
	for _, Rating := range ratings.Ratings {
		User, err := userClient.Get(context.TODO(), &userGw.GetRequest{Id: Rating.UserID})
		if err != nil {
			panic(err)
		}

		accommodationToAdd := domain.RatingDTO{
			UserID:       Rating.UserID,
			TargetId:     Rating.TargetId,
			RatingValue:  int(Rating.RatingValue),
			LastModified: Rating.LastModified,
			Name:         User.User.Name,
			Username:     User.User.Username,
			Surname:      User.User.Surname,
			Id:           Rating.Id,
		}
		ratingsHost = append(ratingsHost, accommodationToAdd)
	}

	responseOne, err := json.Marshal(ratingsHost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseOne)
}
func (handler *UserHandler) ProminentHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hostId := pathParams["id"]
	reservationClient := services.NewReservationClient(handler.ReservationAddress)
	ratingClient := services.NewRatingClient(handler.RatingAddress)
	userClient := services.NewUserClient(handler.UserAddress)
	averageHostRating, err := ratingClient.GetAverageHostRating(context.TODO(), &ratingGw.AverageHostRequest{Id: hostId})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	checkHost, err := reservationClient.CheckReservationRequirementsHost(context.TODO(), &reservationGw.ReservationRequirementsHostRequest{HostId: hostId})
	if averageHostRating.Average >= 4.7 && checkHost.IsPossible {

		_, err = userClient.UpdateProminentStatus(context.TODO(), &userGw.UpdateProminentStatusRequest{Id: hostId, Status: true})
		response, err := json.Marshal(true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response)
	} else {
		_, err = userClient.UpdateProminentStatus(context.TODO(), &userGw.UpdateProminentStatusRequest{Id: hostId, Status: false})
		response, err := json.Marshal(false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response)
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
			if reservation.AccommodationID == idA && reservation.ReservationStatus == "CONFIRMED" {
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
