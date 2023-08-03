package api

import (
	events "github.com/tamararankovic/microservices_demo/common/saga/delete_rating"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_service/application"
	"rating_service/domain"
)

type DeleteRatingCommandHandler struct {
	ratingService     *application.RatingService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeleteRatingCommandHandler(ratingService *application.RatingService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateRatingCommandHandler, error) {
	o := &DeleteRatingCommandHandler{
		ratingService:     ratingService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteRatingCommandHandler) handle(command *events.DeleteRatingCommand) {
	id, err := primitive.ObjectIDFromHex(command.Rating.Id)
	if err != nil {
		return
	}
	order := &domain.Rating{Id: id}

	reply := events.DeleteRatingReply{Rating: command.Rating}

	switch command.Type {
	case events.ShipOrder:
		err := handler.ratingService.Create(order)
		if err != nil {
			reply.Type = events.OrderShippingNotScheduled
			break
		}
		reply.Type = events.OrderShippingScheduled
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
