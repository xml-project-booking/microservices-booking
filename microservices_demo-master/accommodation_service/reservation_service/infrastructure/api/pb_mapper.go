package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	"resevation/domain"
)

func mapReservation(order *domain.Reservation) *pb.Reservation {
	orderPb := &pb.Reservation{
		Id:              order.Id.Hex(),
		AccommodationID: order.AccommodationID.Hex(),
		StartDate:       order.StartDate.Format("2006-01-02"),
		EndDate:         order.EndDate.Format("2006-01-02"),
		GuestNumber:     order.GuestNumber,
		IsConfirmed:     order.Confirmation,
	}
	return orderPb
}

/*func mapReservationCancelation(isDeleted bool) *pb.CancelReservationResponse {
	isDeletedPb := &pb.CancelReservationResponse{
		isDeleted: isDeleted,
	}
}*/
