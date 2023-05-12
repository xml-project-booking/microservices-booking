package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/user_service"
	"user_service/domain"
)

func mapUser(order *domain.User) *pb.User {
	orderPb := &pb.User{
		Id: order.Id.Hex(),
	}
	return orderPb
}
