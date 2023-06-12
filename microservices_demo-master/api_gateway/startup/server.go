package startup

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/api"
	cfg "github.com/tamararankovic/microservices_demo/api_gateway/startup/config"
	reservationGw "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	userGw "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)

	if err != nil {
		panic(err)
	}
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	err1 := reservationGw.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err1 != nil {
		fmt.Printf("dhhdhdhdhdhddhdhdhdhdhdhdhdhdhdhdhdhdh")
		panic(err)
	}

}

func (server *Server) initCustomHandlers() {
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	orderingHandler := api.NewOrderingHandler(orderingEmdpoint, catalogueEmdpoint, shippingEmdpoint)
	orderingHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
