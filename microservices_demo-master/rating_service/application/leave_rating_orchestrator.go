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
		Type:   events.StartedCreatingRating,
		Rating: events.RatingDetails{},
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
	case events.AccommodationUpdate:
		return events.ApproveRating
	case events.AccommodationNotUpdate:
		return events.GetOldValue
	case events.OldValueReturned:
		return events.RollBackRating
	case events.RatingRollBack:
		return events.CancelRating
	default:
		return events.UnknownCommand
	}
}
