package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/api"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	cfg "github.com/tamararankovic/microservices_demo/api_gateway/startup/config"
	reservationGw "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	termGw "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	userGw "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	termEndpoint := fmt.Sprintf("%s:%s", server.config.TermHost, server.config.TermPort)
	err2 := termGw.RegisterTermServiceHandlerFromEndpoint(context.TODO(), server.mux, termEndpoint, opts)
	if err2 != nil {
		fmt.Printf("Problem neki sa term_service u fajlu server.go")
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
	handler := MiddlewareContentTypeSetWithCORS(server.mux)
	newHandler := AuthMiddleware(handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), newHandler))
}
func MiddlewareContentTypeSetWithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		log.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		// Add CORS headers
		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Add Content-Type header
		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if true {
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := authParts[1]
		authenticationEmdpoint := "localhost:8000"
		authenticationClient := services.NewUserClient(authenticationEmdpoint)
		message, _ := authenticationClient.Authenticate(r.Context(), &userGw.AuthenticateRequest{
			Token: token,
		})
		if message.Message == "ok" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized: "+message.Message, http.StatusUnauthorized)
		}
	})
}
