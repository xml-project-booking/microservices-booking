package api

import (
	"context"
	"log"
	"strings"
	"user_service/application"
	"user_service/domain"

	pb "github.com/tamararankovic/microservices_demo/common/proto/user_service"

	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service     *application.UserService
	serviceAuth *application.AuthentificationService
}

func NewUserHandler(service *application.UserService, serviceAuth *application.AuthentificationService) *UserHandler {
	return &UserHandler{
		service:     service,
		serviceAuth: serviceAuth,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	User, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	UserPb := mapUser(User)
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}
<<<<<<< HEAD
func (handler *UserHandler) ExistsUsername(ctx context.Context, request *pb.ExistsUsernameRequest) (*pb.ExistsUsernameResponse, error) {
	username := request.Username
	exists, err := handler.serviceAuth.ExistsUsername(username)
	if err != nil {
		return nil, err
	}
	response := &pb.ExistsUsernameResponse{
		Exists: exists,
	}
	return response, nil
}
func (handler *UserHandler) ExistsEmail(ctx context.Context, request *pb.ExistsEmailRequest) (*pb.ExistsEmailResponse, error) {
	email := request.Email
	exists, err := handler.serviceAuth.ExistsEmail(email)
	if err != nil {
		return nil, err
	}
	response := &pb.ExistsEmailResponse{
		Exists: exists,
=======
func (handler *UserHandler) UpdateCancellationNumber(ctx context.Context, request *pb.UpdateCancellationNumberRequest) (*pb.UpdateCancellationNumberResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	User, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	//UserPb := mapUser(User)
	error := handler.service.UpdateCancellationNumber(User)
	response := &pb.UpdateCancellationNumberResponse{
		CancelNumber: error.Error(),
>>>>>>> 859ba3a (implemented creating of accommodation)
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		current := mapUser(User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}
func (handler *UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := &domain.User{
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Address:  request.Address,
	}

	if request.IsHost {
		user.Role = 0
	} else {
		user.Role = 1
	}

	jwtToken, err := handler.serviceAuth.RegisterUser(user)

	if err != nil {
		return &pb.RegisterResponse{
			RequestResult: &pb.RequestResult{
				Code:    400,
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.RegisterResponse{
		RequestResult: &pb.RequestResult{
			Code: 200,
		},
		Token: jwtToken,
	}, nil
}
func (handler *UserHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("BLAA " + request.Username + request.Password)
	jwtToken, role, id, err := handler.serviceAuth.Login(request.Username, request.Password)

	if err != nil {
		return &pb.LoginResponse{
			RequestResult: &pb.RequestResult{
				Code:    400,
				Message: "Invalid username or password",
			},
		}, nil
	}

	return &pb.LoginResponse{
		Token: jwtToken,
		RequestResult: &pb.RequestResult{
			Code: 200,
		},
		Role: role,
		Id:   id,
	}, nil
}
func (handler *UserHandler) Authenticate(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	log.Println("pozvalo se???")
	message := "ok"

	tokenInfo, err := parseToken(request.Token)
	if err != nil {
		message = "invalid auth token"
	}
	user_id := userClaimFromToken(tokenInfo)
	substrings := strings.Split(user_id, "\"")
	objectId, err := primitive.ObjectIDFromHex(substrings[1])
	if err != nil {
		return nil, err
	}
	user, err := handler.service.Get(objectId)
	if err != nil || user == nil {
		message = "user not found"
	}

	response := &pb.AuthenticateResponse{
		Message: message,
	}
	return response, nil
}
func parseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func userClaimFromToken(claims jwt.MapClaims) string {

	sub, err := claims["user_id"].(string)
	if !err {
		return "error user claim"
	}

	return sub
}
func (handler *UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (response *pb.UpdateUserResponse, err error) {
	user := &domain.User{
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Address:  request.Address,
		Username: request.Username,
		Password: request.Password,
	}
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user.Id = objectId
	err = handler.service.UpdateUser(user)

	if err != nil {
		return &pb.UpdateUserResponse{
			RequestResult: &pb.RequestResult{
				Code:    400,
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.UpdateUserResponse{
		RequestResult: &pb.RequestResult{
			Code: 200,
		},
	}, nil
}
