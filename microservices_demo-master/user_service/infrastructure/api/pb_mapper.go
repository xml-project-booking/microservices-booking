package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"user_service/domain"
)

func mapUser(order *domain.User) *pb.User {
	boolRole := true
	if order.Role == 1 {
		boolRole = false
	}

	orderPb := &pb.User{
		Id:                 order.Id.Hex(),
		Name:               order.Name,
		Surname:            order.Surname,
		Email:              order.Email,
		Address:            order.Address,
		Username:           order.Username,
		IsHost:             boolRole,
		CancellationNumber: int64(order.CancellationNumber),
	}
	return orderPb
}
