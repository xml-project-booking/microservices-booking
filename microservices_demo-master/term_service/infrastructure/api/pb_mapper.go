package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	"term_service/domain"
)

func mapTerm(term *domain.Term) *pb.Term {
	termPb := &pb.Term{
		Id: term.Id.Hex(),
	}
	return termPb
}
