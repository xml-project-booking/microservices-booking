package application

import (
	events "github.com/tamararankovic/microservices_demo/common/saga/leave_rating"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"rating_service/domain"
)

type LeaveRatingOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewLeaveRatingOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*LeaveRatingOrchestrator, error) {
	o := &LeaveRatingOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *LeaveRatingOrchestrator) Start(rating *domain.Rating, oldRating *domain.Rating) error {
	event := &events.LeaveRatingCommand{
		Type: events.StartedCreatingRating,
		Rating: events.RatingDetails{
			ID:           rating.Id,
			TargetID:     rating.TargetId,
			TargetType:   uint32(rating.TargetType),
			UserID:       rating.UserID,
			Value:        uint32(rating.RatingValue),
			LastModified: rating.LastModified,
		},
	}
	if oldRating != nil {
		event.Rating.OldValue = &events.RatingDetails{
			ID:           oldRating.Id,
			TargetID:     oldRating.TargetId,
			TargetType:   uint32(oldRating.TargetType),
			UserID:       oldRating.UserID,
			Value:        uint32(oldRating.RatingValue),
			LastModified: oldRating.LastModified,
		}
	}

	return o.commandPublisher.Publish(event)
}

func (o *LeaveRatingOrchestrator) handle(reply *events.LeaveRatingReply) {
	command := events.LeaveRatingCommand{Rating: reply.Rating}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *LeaveRatingOrchestrator) nextCommandType(reply events.LeaveRatingReplyType) events.LeaveRatingCommandType {
	switch reply {
	case events.CreationStarted:
		return events.CreateRating
	case events.RatingCreated:
		return events.UpdateAccommodation
	case events.AccommodationNotUpdate:
		return events.RollBackRating
	case events.RatingRollBack:
		return events.CancelRating
	case events.AccommodationUpdate:
		return events.SendNotification
	default:
		return events.UnknownCommand
	}
}
