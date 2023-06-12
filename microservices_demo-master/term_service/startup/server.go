package startup

import (
	"fmt"
	term "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"term_service/application"
	"term_service/domain"
	"term_service/infrastructure/api"
	"term_service/infrastructure/persistence"
	"term_service/startup/config"
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
	QueueGroup = "term_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	termStore := server.initTermStore(mongoClient)

	termService := server.initTermService(termStore)

	termHandler := server.initTermHandler(termService)

	server.startGrpcServer(termHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	fmt.Print(server.config.TermDBHost, server.config.TermDBPort)
	client, err := persistence.GetClient(server.config.TermDBHost, server.config.TermDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initTermStore(client *mongo.Client) domain.TermStore {
	store := persistence.NewTermMongoDBStore(client)
	store.DeleteAll()
	for _, Term := range terms {
		err := store.Insert(Term)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

//func (server *Server) initPublisher(subject string) saga.Publisher {
//	publisher, err := nats.NewNATSPublisher(
//		server.config.NatsHost, server.config.NatsPort,
//		server.config.NatsUser, server.config.NatsPass, subject)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return publisher
//}
//
//func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
//	subscriber, err := nats.NewNATSSubscriber(
//		server.config.NatsHost, server.config.NatsPort,
//		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return subscriber
//}

func (server *Server) initTermService(store domain.TermStore) *application.TermService {
	return application.NewTermService(store)
}

//func (server *Server) initCreateOrderHandler(service *application.TermService, publisher saga.Publisher, subscriber saga.Subscriber) {
//	_, err := api.NewCreateTermCommandHandler(service, publisher, subscriber)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func (server *Server) initTermHandler(service *application.TermService) *api.TermHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.TermHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	term.RegisterTermServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
