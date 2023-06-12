package middleware

// import (
// 	"context"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"strings"
// 	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"

// )

// type JwtContent struct{}

// func MiddlewareAuthentification(next http.Handler, authService *authentification.JwtService) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
// 		jwt_bearer := h.Header.Get("Authorization")

// 		parts := strings.Split(jwt_bearer, "Bearer ")

// 		if len(parts) < 2 {
// 			http.Error(rw, "Token not present", http.StatusUnauthorized)
// 			return
// 		}

// 		jwt := parts[1]

// 		if err != nil {
// 			http.Error(rw, "Something went wrong", http.StatusInternalServerError)
// 			return
// 		}

// 		claims, result := authService.ValidateToken(jwt, key.PublicKey)

// 		if result == authentification.Token_Expired {
// 			http.Error(rw, "Token expired", http.StatusUnauthorized)
// 			return
// 		}

// 		if result == authentification.Token_Invalid {
// 			http.Error(rw, "Token invalid", http.StatusUnauthorized)
// 			return
// 		}

// 		if result == authentification.Token_Failed {
// 			http.Error(rw, "Something failed with token", http.StatusInternalServerError)
// 			return
// 		}

// 		ctx := context.WithValue(h.Context(), JwtContent{}, claims)
// 		h = h.WithContext(ctx)

// 		next.ServeHTTP(rw, h)
// 	})
// }

// func MiddlewareContentTypeSet(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
// 		log.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

// 		rw.Header().Add("Content-Type", "application/json")

// 		next.ServeHTTP(rw, h)
// 	})
// }

// func MiddlewareAuthorization(next http.Handler, roles []string) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
// 		jwt_claims := h.Context().Value(JwtContent{}).(*authentification.SignedDetails)

// 		if !slices.Contains(roles, jwt_claims.Role) {
// 			http.Error(rw, "You are not authorized", http.StatusBadRequest)
// 			return
// 		}

// 		next.ServeHTTP(rw, h)
// 	})
// }
// func MiddlewareContentTypeSetWithCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
// 		log.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

// 		// Add CORS headers
// 		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
// 		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		// Add Content-Type header
// 		rw.Header().Add("Content-Type", "application/json")

// 		next.ServeHTTP(rw, h)
// 	})
// }
// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if true {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		authParts := strings.Split(authHeader, " ")
// 		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		token := authParts[1]
// 		authenticationEmdpoint := "localhost:8000"
// 		authenticationClient := services.NewUserClient(authenticationEmdpoint)
// 		message, _ := authenticationClient.Authenticate(r.Context(), &authentication_Gw.AuthenticateRequest{
// 			Token: token,
// 		})
// 		if message.Message == "ok" {
// 			next.ServeHTTP(w, r)
// 		} else {
// 			http.Error(w, "Unauthorized: "+message.Message, http.StatusUnauthorized)
// 		}
// 	})
// }
