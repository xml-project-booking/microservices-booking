package services

import (
	accommodation "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	catalogue "github.com/tamararankovic/microservices_demo/common/proto/catalogue_service"
	ordering "github.com/tamararankovic/microservices_demo/common/proto/ordering_service"
	ratings "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	reservation "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	shipping "github.com/tamararankovic/microservices_demo/common/proto/shipping_service"
	terms "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	user "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewCatalogueClient(address string) catalogue.CatalogueServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return catalogue.NewCatalogueServiceClient(conn)
}

func NewOrderingClient(address string) ordering.OrderingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Ordering service: %v", err)
	}
	return ordering.NewOrderingServiceClient(conn)
}

func NewShippingClient(address string) shipping.ShippingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return shipping.NewShippingServiceClient(conn)
}
func NewUserClient(address string) user.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return user.NewUserServiceClient(conn)
}
func NewReservationClient(address string) reservation.ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Reservation service: %v", err)
	}
	return reservation.NewReservationServiceClient(conn)
}
func NewAccommodationClient(address string) accommodation.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation.NewAccommodationServiceClient(conn)
}
func NewTermClient(address string) terms.TermServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return terms.NewTermServiceClient(conn)

}
func NewRatingClient(address string) ratings.RatingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return ratings.NewRatingServiceClient(conn)

}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(100*1024*1024)))
}
