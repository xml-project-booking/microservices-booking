package startup

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/api"
	cfg "github.com/tamararankovic/microservices_demo/api_gateway/startup/config"
	accommodationGw "github.com/tamararankovic/microservices_demo/common/proto/accommodation_service"
	ratingGw "github.com/tamararankovic/microservices_demo/common/proto/rating_service"
	reservationGw "github.com/tamararankovic/microservices_demo/common/proto/reservation_service"
	termGw "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	userGw "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024), grpc.MaxCallSendMsgSize(100*1024*1024)), // Set maximum frame size to 10MB
	}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)

	if err != nil {
		panic(err)
	}
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	err1 := reservationGw.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), server.mux, reservationEndpoint, opts)
	if err1 != nil {
		panic(err)
	}

	termEndpoint := fmt.Sprintf("%s:%s", server.config.TermHost, server.config.TermPort)
	err2 := termGw.RegisterTermServiceHandlerFromEndpoint(context.TODO(), server.mux, termEndpoint, opts)
	if err2 != nil {
		fmt.Printf("Problem neki sa term_service u fajlu server.go")
		panic(err)
	}

	accommodationPoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err3 := accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationPoint, opts)
	if err3 != nil {
		panic(err)
	}
	ratingPoint := fmt.Sprintf("%s:%s", server.config.RatingHost, server.config.RatingPort)
	err4 := ratingGw.RegisterRatingServiceHandlerFromEndpoint(context.TODO(), server.mux, ratingPoint, opts)
	if err4 != nil {
		fmt.Printf("dhhdhdhdhdhddhdhdhdhdhdhdhdhdhdhdhdhdh")
		panic(err)
	}

}

func (server *Server) initCustomHandlers() {
	reservationEndpoint := fmt.Sprintf("%s:%s", server.config.ReservationHost, server.config.ReservationPort)
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	termEndpoint := fmt.Sprintf("%s:%s", server.config.TermHost, server.config.TermPort)
	ratingEndpoint := fmt.Sprintf("%s:%s", server.config.RatingHost, server.config.RatingPort)
	reservationHandler := api.NewReservationHandler(reservationEndpoint, userEndpoint, accommodationEndpoint, termEndpoint)
	accommodationHandler := api.NewAccommodationHandler(reservationEndpoint, userEndpoint, accommodationEndpoint, termEndpoint)
	reservationHandler.Init(server.mux)
	accommodationHandler.Init(server.mux)

	//delete-user
	userHandler := api.NewUserHandler(userEndpoint, reservationEndpoint, accommodationEndpoint, ratingEndpoint)
	userHandler.Init(server.mux)
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
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, Accept-Encoding, X-CSRF-Token, accept, origin, Cache-Control, X-Requested-With")

		// Add Content-Type header
		rw.Header().Add("Content-Type", "application/json")
		if h.Method == "OPTIONS" {
			// Handle preflight request
			rw.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(rw, h)
	})
}
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*if r.URL.Path == "/users/register" || r.URL.Path == "/users/login" || strings.Contains(r.URL.Path, "/users/user/existsUsername") || strings.Contains(r.URL.Path, "/users/user/existsEmail") || strings.Contains(r.URL.Path, "/users/authenticate") {
			// Call the next handler without performing authentication
			next.ServeHTTP(w, r)
			return
		}
		if r.Method == http.MethodOptions {
			// Set the necessary CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, Accept-Encoding, X-CSRF-Token, accept, origin, Cache-Control, X-Requested-With")
			w.WriteHeader(http.StatusOK)
			return
		}*/
		//otkom ovo i zakom ovo iznad ako testiratenesto bez aut
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
		log.Println(token)
		//authenticationEndpoint := "localhost:8000"
		//authenticationClient := services.NewUserClient(authenticationEndpoint)
		//message, err := authenticationClient.Authenticate(r.Context(), &userGw.AuthenticateRequest{
		//	Token: token,
		//})
		authServiceEndpoint := "localhost:8000"
		resp, err := http.Get("http://" + authServiceEndpoint + "/users/authenticate/" + token)
		if err != nil {
			log.Println("Error during authentication:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		responseContent := string(body)
		var response map[string]interface{}
		err2 := json.Unmarshal([]byte(responseContent), &response)
		if err2 != nil {
			log.Println("Failed to decode response:", err2)
			return
		}

		messageValue, ok := response["message"].(string)
		if !ok {
			log.Println("Failed to get message value")
			return
		}
		if messageValue == "ok" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized: "+messageValue, http.StatusUnauthorized)
		}
	})
}
