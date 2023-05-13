package api

import (
	pb "/common/proto/term_service"
	"context"
	"term_service/application"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TermHandler struct {
	pb.UnimplementedTermServiceServer
	service *application.TermService
}

func NewUserHandler(service *application.TermService) *TermHandler {
	return &TermHandler{
		service: service,
	}
}

func (handler *TermHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Term, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	TermPb := mapTerm(Term)
	response := &pb.GetResponse{
		Term: TermPb,
	}
	return response, nil
}

func (handler *TermHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Terms, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Terms: []*pb.Term{},
	}
	for _, Term := range Terms {
		current := mapTerm(Term)
		response.Terms = append(response.Terms, current)
	}
	return response, nil
}
