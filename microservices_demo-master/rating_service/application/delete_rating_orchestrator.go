package application

import (
	"github.com/tamararankovic/microservices_demo/common/saga/delete_rating"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_service/domain"
)

type DeleteRatingOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewDeleteRatingOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteRatingOrchestrator, error) {
	o := &DeleteRatingOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *DeleteRatingOrchestrator) Start(id primitive.ObjectID, oldValue *domain.Rating) error {
	event := &delete_rating.DeleteRatingCommand{
		Type:   delete_rating.StartedDeletionRating,
		Rating: delete_rating.RatingDetails{},
	}
	return o.commandPublisher.Publish(event)
}

func (o *DeleteRatingOrchestrator) handle(reply *delete_rating.DeleteRatingReply) {
	command := delete_rating.DeleteRatingCommand{Rating: reply.Rating}
	command.Type = o.nextCommandType(reply)
	if command.Type != delete_rating.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeleteRatingOrchestrator) nextCommandType(reply *delete_rating.DeleteRatingReply) delete_rating.DeleteRatingCommandType {
	switch reply.Type {
	case delete_rating.DeletionStarted:
		if reply.Rating.TargetType == 1 {
			return delete_rating.UpdateHost
		}
		return delete_rating.UpdateAccommodation
	case delete_rating.DeletionFailed:
		return delete_rating.CancelDeletionRating
	case delete_rating.AccommodationUpdate:
		return delete_rating.FinishDeletionRating
	case delete_rating.AccommodationNotUpdate:
		return delete_rating.RollbackRating
	case delete_rating.HostUpdate:
		return delete_rating.FinishDeletionRating
	case delete_rating.HostNotUpdate:
		return delete_rating.RollbackRating
	default:
		return delete_rating.UnknownCommand
	}
}
