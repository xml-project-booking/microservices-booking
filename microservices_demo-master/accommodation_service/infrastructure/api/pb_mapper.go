package api

import (
	"accommodation_service/domain"
	pb "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
)

func mapAccommodation(order *domain.Accommodation) *pb.Accommodation {
	orderPb := &pb.Accommodation{
		Id:                           order.Id.Hex(),
		AccommodationReservationType: order.ReservationConfirmation,
		MinGuest:                     int64(order.MinGuest),
		MaxGuest:                     int64(order.MaxGuest),
		Name:                         order.Name,
		Country:                      order.Country,
		City:                         order.City,
		Street:                       order.Street,
		Wifi:                         order.Wifi,
		Kitchen:                      order.Kitchen,
		FreeParking:                  order.FreeParking,
		AirConditioning:              order.AirConditioning,
	}
	return orderPb
}
