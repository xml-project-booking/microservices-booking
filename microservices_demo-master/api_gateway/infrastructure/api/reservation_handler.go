package api

import (
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	accommodations "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	reservations "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	terms "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	users "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type ReservationHandler struct {
	reservationClientAddress string
	userClientAddress        string
	accommodationAddress     string
	termAddress              string
}

func (handler *ReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/reservations-cancel/{reservationId}/{userId}", handler.CancelReservation)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/reservations-host/{accommodationId}", handler.GetReservationForHost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/reservations-host-confirmed/{accommodationId}", handler.GetReservationForHostConfirmed)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/make-reservation", handler.ConfirmationOfReservation)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/confirm-reservation-man/{reservationId}", handler.ManuallyConfirmReservation)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/cancel-reservation-man/{reservationId}", handler.ManuallyCancelReservation)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/term/reservation", handler.CheckReservationForTerm)
	if err != nil {
		panic(err)
	}
}

func NewReservationHandler(reservationClientAddress, userClientAddress, accommodationAddress, termAddress string) Handler {
	return &ReservationHandler{
		reservationClientAddress: reservationClientAddress,
		userClientAddress:        userClientAddress,
		accommodationAddress:     accommodationAddress,
		termAddress:              termAddress,
	}
}

func (handler *ReservationHandler) CancelReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userIdString := pathParams["userId"]
	reservationIdString := pathParams["reservationId"]
	fmt.Println(userIdString + "fhfhhfhf")
	fmt.Println(reservationIdString + "ovo je id")
	fmt.Println("usaoo u pogresnu funkciju")
	userId, err := primitive.ObjectIDFromHex(userIdString)
	reservationId, err := primitive.ObjectIDFromHex(reservationIdString)
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	userClient := services.NewUserClient(handler.userClientAddress)
	if userIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if reservationIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//cancelReservation := &domain.CancelReservation{UserId: userId,
	//ReservationId: reservationId}
	cancelNumber, err := userClient.UpdateCancellationNumber(context.TODO(), &users.UpdateCancellationNumberRequest{
		Id: userId.Hex(),
	})
	fmt.Println(cancelNumber)
	cancellationReservation, error := reservationClient.CancelReservation(context.TODO(), &reservations.CancelReservationRequest{Id: reservationId.Hex()})
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//handler.addShippingInfo(orderDetails)
	//handler.addProductInfo(orderDetails)

	response, err := json.Marshal(cancellationReservation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func (handler *ReservationHandler) GetReservationForHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	accommodationId := pathParams["accommodationId"]
	objectId, err := primitive.ObjectIDFromHex(accommodationId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if accommodationId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reservationList := handler.GetReservations(objectId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(reservationList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *ReservationHandler) GetReservationForHostConfirmed(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	accommodationId := pathParams["accommodationId"]
	objectId, err := primitive.ObjectIDFromHex(accommodationId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if accommodationId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reservationList := handler.GetReservationsConfirmed(objectId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(reservationList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
func (handler *ReservationHandler) GetReservationsConfirmed(accommodationId primitive.ObjectID) []*domain.ReservationForHost {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	userClient := services.NewUserClient(handler.userClientAddress)
	reservationsAccommodation, err := reservationClient.GetAllByAccommodationConfirmed(context.TODO(), &reservations.GetAllByAccommodationRequest{Id: accommodationId.Hex()})
	if err != nil {

		panic(err)
	}
	var reservationSlice []*domain.ReservationForHost
	for _, reservation := range reservationsAccommodation.Reservations {
		user, err := userClient.Get(context.TODO(), &users.GetRequest{
			Id: reservation.GuestId,
		},
		)
		if err != nil {
			panic(err)
		}

		reservationDetails := &domain.ReservationForHost{
			Id:                 reservation.Id,
			StartDate:          reservation.StartDate,
			EndDate:            reservation.EndDate,
			GuestNumber:        reservation.GuestNumber,
			CancellationNumber: user.User.CancellationNumber,
			GuestName:          user.User.Name,
			GuestSurname:       user.User.Surname,
			GuestId:            user.User.Id,
		}
		reservationSlice = append(reservationSlice, reservationDetails)
	}
	fmt.Println(len(reservationSlice))
	fmt.Println("ovo je duzina")
	return reservationSlice
}
func (handler *ReservationHandler) GetReservations(accommodationId primitive.ObjectID) []*domain.ReservationForHost {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	userClient := services.NewUserClient(handler.userClientAddress)
	reservationsAccommodation, err := reservationClient.GetAllByAccommodation(context.TODO(), &reservations.GetAllByAccommodationRequest{Id: accommodationId.Hex()})
	if err != nil {

		panic(err)
	}
	var reservationSlice []*domain.ReservationForHost
	for _, reservation := range reservationsAccommodation.Reservations {
		user, err := userClient.Get(context.TODO(), &users.GetRequest{
			Id: reservation.GuestId,
		},
		)
		if err != nil {
			panic(err)
		}

		reservationDetails := &domain.ReservationForHost{
			Id:                 reservation.Id,
			StartDate:          reservation.StartDate,
			EndDate:            reservation.EndDate,
			GuestNumber:        reservation.GuestNumber,
			CancellationNumber: user.User.CancellationNumber,
			GuestName:          user.User.Name,
			GuestSurname:       user.User.Surname,
			GuestId:            user.User.Id,
		}
		reservationSlice = append(reservationSlice, reservationDetails)
	}
	fmt.Println(len(reservationSlice))
	fmt.Println("ovo je duzina")
	return reservationSlice
}
func (handler *ReservationHandler) ConfirmationOfReservation(w http.ResponseWriter, req *http.Request, params map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationAddress)
	//termClient := services.NewTermClient(handler.termAddress)

	decoder := json.NewDecoder(req.Body)
	var t domain.ReservationDTO
	err := decoder.Decode(&t)
	fmt.Println(req.Body)
	if err != nil {
		return

	}
	accommodation, err := accommodationClient.Get(context.TODO(), &accommodations.GetRequest{
		Id: t.AccommodationID,
	})
	fmt.Println(accommodation)
	fmt.Println(accommodation.Accommodation.AccommodationReservationType + "ovo je tip")
	if accommodation.Accommodation.AccommodationReservationType == "AUTOMATIC" {
		fmt.Println("usaooooooooo u automatic")
		res, err := reservationClient.ConfirmReservationAutomatically(context.TODO(), &reservations.ReservationRequest{
			Id:              t.Id,
			AccommodationID: t.AccommodationID,
			StartDate:       t.StartDate,
			EndDate:         t.EndDate,
			GuestNumber:     t.GuestNumber,
			GuestId:         t.GuestId,
			Confirmation:    t.Confirmation,
			MinGuest:        strconv.FormatInt(accommodation.Accommodation.MinGuest, 10),
			MaxGuest:        strconv.FormatInt(accommodation.Accommodation.MaxGuest, 10),
			HostId:          t.HostId,
		})
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id := res.Id
		response, err := json.Marshal(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write((response))
		return
	}

	fmt.Println("usaoooooooo u manuall")
	result, err := reservationClient.MakeRequestForReservation(context.TODO(), &reservations.ReservationRequest{Id: t.Id,
		AccommodationID: t.AccommodationID,
		StartDate:       t.StartDate,
		EndDate:         t.EndDate,
		GuestNumber:     t.GuestNumber,
		GuestId:         t.GuestId,
		Confirmation:    t.Confirmation,
		MinGuest:        strconv.FormatInt(accommodation.Accommodation.MinGuest, 10),
		MaxGuest:        strconv.FormatInt(accommodation.Accommodation.MaxGuest, 10),
		HostId:          t.HostId},
	)

	if err != nil {
		return
	}
	fmt.Println(t)
	fmt.Println("hahahahahahahahahahahahahhaha")
	id := result.Id
	response, err := json.Marshal(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write((response))
	return
}

func (handler *ReservationHandler) ManuallyConfirmReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	id := pathParams["reservationId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := reservationClient.ConfirmReservationManually(context.TODO(), &reservations.ConfirmReservationManuallyRequest{Id: id})
	if err != nil {
		return
	}
	idResult := result.Id
	response, err := json.Marshal(idResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write((response))

	return
}
func (handler *ReservationHandler) ManuallyCancelReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	id := pathParams["reservationId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := reservationClient.CancelReservationManually(context.TODO(), &reservations.CancelReservationManuallyRequest{Id: id})
	if err != nil {
		return
	}
	idResult := result.Id
	response, err := json.Marshal(idResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write((response))

	return
}

func (handler *ReservationHandler) CheckReservationForTerm(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	termClient := services.NewTermClient(handler.termAddress)
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	decoder := json.NewDecoder(r.Body)
	var t domain.TermDTO
	err := decoder.Decode(&t)
	if err != nil {
		return
	}
	fmt.Println(t)
	res, err := reservationClient.TermCheck(context.TODO(), &reservations.TermCheckRequest{
		Id:        t.AccommodationID,
		Startdate: t.StartDate,
		EndDate:   t.EndDate,
	})

	if res.HasReservation != "greska" {
		w.WriteHeader(http.StatusForbidden)
		return

	}
	num, err := strconv.Atoi(t.Value)
	termClient.Create(context.TODO(), &terms.CreateRequest{
		UserId:          t.UserID,
		StartDate:       t.StartDate,
		EndDate:         t.EndDate,
		AccommodationId: t.AccommodationID,
		PriceType:       t.PriceType,
		Value:           float64(num),
	})
	w.WriteHeader(http.StatusOK)
	return
}
