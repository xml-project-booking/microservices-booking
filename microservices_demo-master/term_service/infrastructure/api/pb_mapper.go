package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/term_service"
	"term_service/domain"
)

func mapTerm(term *domain.Term) *pb.GetResponse {
	termPb := &pb.GetResponse{
		Id:              term.Id.Hex(),
		Date:            term.Date.String(),
		PriceType:       term.PriceType,
		Value:           float64(term.Value),
		AccommodationId: term.AccommodationID.Hex(),
		UserId:          term.UserID.Hex(),
	}
	return termPb
}

func mapTermOne(term *domain.Term) *pb.Term {
	termPb := &pb.Term{
		Id:              term.Id.Hex(),
		Date:            term.Date.String(),
		PriceType:       term.PriceType,
		Value:           float64(term.Value),
		AccommodationId: term.AccommodationID.Hex(),
		UserId:          term.UserID.Hex(),
	}
	return termPb
}
