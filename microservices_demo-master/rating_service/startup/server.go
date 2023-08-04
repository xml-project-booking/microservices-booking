package startup

import (
	"fmt"
	rating "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"github.com/tamararankovic/microservices_demo/common/saga/messaging/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"rating_service/application"
	"rating_service/domain"
	"rating_service/infrastructure/api"
	"rating_service/infrastructure/persistence"
	"rating_service/startup/config"
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
	QueueGroup = "rating_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	ratingStore := server.initRatingStore(mongoClient)

	ratingService := server.initRatingService(ratingStore)

	ratingHandler := server.initRatingHandler(ratingService)

	server.startGrpcServer(ratingHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	fmt.Print(server.config.RatingDBHost, server.config.RatingDBPort)
	client, err := persistence.GetClient(server.config.RatingDBHost, server.config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initRatingStore(client *mongo.Client) domain.RatingStore {
	store := persistence.NewRatingMongoDBStore(client)
	store.DeleteAll()
	for _, Rating := range ratings {
		err := store.Insert(Rating)
		if err != nil {
			log.Fatal(err)
		}
	}
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

func (server *Server) initRatingService(store domain.RatingStore) *application.RatingService {
	return application.NewRatingService(store)
}

func (server *Server) initLeaveRatingHandler(service *application.RatingService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewLeaveRatingCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initRatingHandler(service *application.RatingService) *api.RatingHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(ratingHandler *api.RatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, ratingHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
