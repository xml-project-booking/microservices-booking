package api

import (
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	accommodations "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	reservations "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	users "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"net/http"
)

type ReservationHandler struct {
	reservationClientAddress string
	userClientAddress        string
	accommodationAddress     string
}

func (handler *ReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/reservations/{reservationId}/{userId}", handler.CancelReservation)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/reservations-host/{accommodationId}", handler.GetReservationForHost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/make-reservation", handler.GetReservationForHost)
	if err != nil {
		panic(err)
	}
}

func NewReservationHandler(reservationClientAddress, userClientAddress, accommodationAddress string) Handler {
	return &ReservationHandler{
		reservationClientAddress: reservationClientAddress,
		userClientAddress:        userClientAddress,
		accommodationAddress:     accommodationAddress,
	}
}

func (handler *ReservationHandler) CancelReservation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	userIdString := pathParams["userId"]
	reservationIdString := pathParams["reservationId"]
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
func (handler *ReservationHandler) GetReservations(accommodationId primitive.ObjectID) []*domain.ReservationForHost {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	userClient := services.NewUserClient(handler.userClientAddress)
	reservationsAccommodation, err := reservationClient.GetAllByAccommodation(context.TODO(), &reservations.GetAllByAccommodationRequest{Id: accommodationId.Hex()})
	if err != nil {

		return nil
	}
	var reservationSlice []*domain.ReservationForHost
	for _, reservation := range reservationsAccommodation.Reservations {
		user, err := userClient.Get(context.TODO(), &users.GetRequest{
			Id: reservation.GuestId,
		})
		if err != nil {
			return nil
		}

		reservationDetails := &domain.ReservationForHost{
			Id:                 reservation.Id,
			StartDate:          reservation.StartDate,
			EndDate:            reservation.EndDate,
			GuestNumber:        reservation.GuestNumber,
			CancellationNumber: user.User.CancellationNumber,
		}
		reservationSlice = append(reservationSlice, reservationDetails)
	}
	return reservationSlice
}
func (handler *ReservationHandler) ConfirmationOfReservation(w http.ResponseWriter, req *http.Request) {
	//reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationAddress)

	decoder := json.NewDecoder(req.Body)
	var t domain.ReservationDTO
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)

	}
	accommodation, err := accommodationClient.Get(context.TODO(), &accommodations.GetRequest{
		Id: t.AccommodationID,
	})
	if accommodation.Accommodation.AccommodationReservationType == "AUTOMATIC" {
		//response, err:= reservationClient.ConfirmReservationAutomatically(context.TODO(),&reservations.ReservationRequest{})
	}
	{
		//response,err:= reservationClient.MakeRequestForReservation(context.TODO(),&reservations.ReservationRequest{});
	}

	fmt.Println(t)
	fmt.Println("hahahahahahahahahahahahahhaha")
}

/*func (handler *ReservationHandler) ManuallyConfirmReservation(orderDetails *domain.OrderDetails) error {
	orderingClient := services.NewOrderingClient(handler.orderingClientAddress)
	orderInfo, err := orderingClient.Get(context.TODO(), &ordering.GetRequest{Id: orderDetails.Id})
	if err != nil {
		return err
	}
	orderDetails.Id = orderInfo.Order.Id
	orderDetails.CreatedAt = orderInfo.Order.CreatedAt.AsTime()
	orderDetails.Status = orderInfo.Order.Status.String()
	orderDetails.OrderItems = make([]domain.OrderItem, 0)
	for _, item := range orderInfo.Order.Items {
		itemDetails := domain.OrderItem{
			Product:  domain.Product{Id: item.Product.Id, ColorCode: item.Product.Color.Code},
			Quantity: uint16(item.Quantity),
		}
		orderDetails.OrderItems = append(orderDetails.OrderItems, itemDetails)
	}
	return nil
}*/
