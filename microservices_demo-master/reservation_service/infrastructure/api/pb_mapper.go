package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	"resevation/domain"
)

func mapReservation(reservation *domain.Reservation) *pb.Reservation {
	orderPb := &pb.Reservation{
		Id:          reservation.Id.Hex(),
		StartDate:   reservation.StartDate.String(),
		EndDate:     reservation.EndDate.String(),
		GuestNumber: reservation.GuestNumber,
		GuestId:     reservation.GuestId.Hex(),
	}
	return orderPb
}

/*func mapReservationCancelation(isDeleted bool) *pb.CancelReservationResponse {
	isDeletedPb := &pb.CancelReservationResponse{
		isDeleted: isDeleted,
	}
}*/
