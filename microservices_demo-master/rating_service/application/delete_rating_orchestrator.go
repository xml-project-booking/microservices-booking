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
		Type: delete_rating.StartedDeletionRating,
		Rating: delete_rating.RatingDetails{
			ID: id,
			OldValue: &delete_rating.RatingDetails{
				ID:           oldValue.Id,
				TargetID:     oldValue.TargetId,
				TargetType:   uint32(oldValue.TargetType),
				UserID:       oldValue.UserID,
				Value:        uint32(oldValue.RatingValue),
				LastModified: oldValue.LastModified,
			},
		},
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
		return delete_rating.DeleteRating
	case delete_rating.RatingDelete:
		return delete_rating.SendNotification

	default:
		return delete_rating.UnknownCommand
	}
}
