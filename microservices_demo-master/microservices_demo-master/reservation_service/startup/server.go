package startup

import (
	"fmt"
	reservation "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"

	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"github.com/tamararankovic/microservices_demo/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"resevation/application"
	"resevation/domain"
	"resevation/infrastructure/api"
	"resevation/infrastructure/persistence"
	"resevation/startup/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	reservationStore := server.initReservationStore(mongoClient)

	reservationService := server.initReservationService(reservationStore)

	reservationHandler := server.initReservationHandler(reservationService)

	server.startGrpcServer(reservationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ReservationDBHost, server.config.ReservationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initReservationStore(client *mongo.Client) domain.ReservationStore {
	store := persistence.NewReservationMongoDBStore(client)
	store.DeleteAll()

	return store
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initReservationService(store domain.ReservationStore) *application.ReservationService {
	return application.NewReservationService(store)
}

func (server *Server) initReservationHandler(service *application.ReservationService) *api.ReservationHandler {
	return api.NewReservationHandler(service)
}

func (server *Server) startGrpcServer(reservationHandler *api.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reservation.RegisterReservationServiceServer(grpcServer, reservationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
