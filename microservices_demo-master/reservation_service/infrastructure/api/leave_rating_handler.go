package api

import (
	"fmt"
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
	fmt.Println(o.commandSubscriber)
	return o, nil
}

func (handler *LeaveRatingCommandHandler) handle(command *events.LeaveRatingCommand) {
	var canLeaveRating = false
	reply := events.LeaveRatingReply{Rating: command.Rating}
	switch command.Type {
	case events.StartedCreatingRating:
		//var canLeaveRating = false
		if command.Rating.TargetType == 0 {
			fmt.Println("2")
			canLeaveRating = handler.reservationService.CheckGuestCanLeaveRating(command.Rating.TargetID, command.Rating.UserID)
		} else {
			canLeaveRating = handler.reservationService.CheckGuestCanLeaveRatingForHost(command.Rating.TargetID, command.Rating.UserID)
		}
		fmt.Println(canLeaveRating)
		if canLeaveRating {
			fmt.Println("8888")
			reply.Type = events.CreationStarted
			break
		}
		reply.Type = events.CantGiveRating
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
