package api

import (
	"accommodation_service/domain"
	pb "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func mapAccommodationPb(order *pb.Accommodation) *domain.Accommodation {
	orderDomain := &domain.Accommodation{
		Id:                      getObjectId(order.Id),
		ReservationConfirmation: order.AccommodationReservationType,
		MinGuest:                int(order.MinGuest),
		MaxGuest:                int(order.MaxGuest),
		Name:                    order.Name,
		Country:                 order.Country,
		City:                    order.City,
		Street:                  order.Street,
		Wifi:                    order.Wifi,
		Kitchen:                 order.Kitchen,
		FreeParking:             order.FreeParking,
		AirConditioning:         order.AirConditioning,
	}
	return orderDomain
}
func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
