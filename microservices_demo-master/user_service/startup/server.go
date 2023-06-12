package startup

import (
	"fmt"
	user "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"user_service/application"
	"user_service/domain"
	"user_service/infrastructure/api"
	"user_service/infrastructure/persistence"
	"user_service/startup/config"
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
	userStore := server.initUserStore(mongoClient)

	userService := server.initUserService(userStore)
	authService := server.initAuthService(userStore)
	userHandler := server.initUserHandler(userService, authService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	fmt.Print(server.config.UserDBHost, server.config.UserDBPort)
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll()
	for _, User := range users {
		err := store.Insert(User)
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

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}
func (server *Server) initAuthService(store domain.UserStore) *application.AuthentificationService {
	return application.NewAuthentificationService(store)
}

//func (server *Server) initCreateOrderHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
//	_, err := api.NewCreateUserCommandHandler(service, publisher, subscriber)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func (server *Server) initUserHandler(service *application.UserService, auth *application.AuthentificationService) *api.UserHandler {
	return api.NewUserHandler(service, auth)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	user.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
