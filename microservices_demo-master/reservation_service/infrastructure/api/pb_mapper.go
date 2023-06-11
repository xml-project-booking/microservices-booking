package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	"resevation/domain"
)

func mapReservation(order *domain.Reservation) *pb.Reservation {
	orderPb := &pb.Reservation{
		Id: order.Id.Hex(),
	}
	return orderPb
}

/*func mapReservationCancelation(isDeleted bool) *pb.CancelReservationResponse {
	isDeletedPb := &pb.CancelReservationResponse{
		isDeleted: isDeleted,
	}
}*/
