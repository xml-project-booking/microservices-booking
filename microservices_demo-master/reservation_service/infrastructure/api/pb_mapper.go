package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	"resevation/domain"
)

func mapReservation(reservation *domain.Reservation) *pb.Reservation {
	orderPb := &pb.Reservation{
		Id:                reservation.Id.Hex(),
		AccommodationID:   reservation.AccommodationID.Hex(),
		StartDate:         reservation.StartDate.Format("2006-01-02"),
		EndDate:           reservation.EndDate.Format("2006-01-02"),
		GuestNumber:       reservation.GuestNumber,
		GuestId:           reservation.GuestId.Hex(),
		IsConfirmed:       reservation.Confirmation,
		ReservationStatus: reservation.ReservationStatus,
	}
	return orderPb
}

/*func mapReservationCancelation(isDeleted bool) *pb.CancelReservationResponse {
	isDeletedPb := &pb.CancelReservationResponse{
		isDeleted: isDeleted,
	}
}*/
