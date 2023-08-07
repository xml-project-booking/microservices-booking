package api

import (
	events "github.com/tamararankovic/microservices_demo/common/saga/leave_rating"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"

	"resevation/application"
)

type LeaveRatingCommandHandler struct {
	reservationService *application.ReservationService
	replyPublisher     saga.Publisher
	commandSubscriber  saga.Subscriber
}

func NewLeaveRatingCommandHandler(reservationService *application.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber) (*LeaveRatingCommandHandler, error) {
	o := &LeaveRatingCommandHandler{
		reservationService: reservationService,
		replyPublisher:     publisher,
		commandSubscriber:  subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *LeaveRatingCommandHandler) handle(command *events.LeaveRatingCommand) {
	reply := events.LeaveRatingReply{Rating: command.Rating}
	switch command.Type {
	case events.StartedCreatingRating:
		canLeaveRating := handler.reservationService.CheckGuestCanLeaveRating(command.Rating.TargetID, command.Rating.UserID)
		if canLeaveRating {
			reply.Type = events.CreationStarted
			break
		}
		reply.Type = events.CreationFailed
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
