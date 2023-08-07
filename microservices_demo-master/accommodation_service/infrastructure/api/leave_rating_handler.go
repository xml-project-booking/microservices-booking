package api

import (
	"accommodation_service/application"
	"accommodation_service/domain"
	events "github.com/tamararankovic/microservices_demo/common/saga/leave_rating"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"log"
)

type LeaveRatingCommandHandler struct {
	ratingService        *application.RatingService
	accommodationService *application.AccommodationService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
}

func NewLeaveRatingCommandHandler(ratingService *application.RatingService, publisher saga.Publisher, subscriber saga.Subscriber, accommodationService *application.AccommodationService) (*LeaveRatingCommandHandler, error) {
	o := &LeaveRatingCommandHandler{
		ratingService:        ratingService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
		accommodationService: accommodationService,
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
	case events.UpdateAccommodation:

		var err error
		r := domain.Rating{
			Id:           command.Rating.ID,
			UserID:       command.Rating.UserID,
			TargetId:     command.Rating.TargetID,
			RatingValue:  int32(command.Rating.Value),
			TargetType:   int(command.Rating.TargetType),
			LastModified: command.Rating.LastModified,
		}
		Accommodation, err := handler.accommodationService.Get(r.TargetId)
		average, err := handler.ratingService.GetAccommodationAverage(r.TargetId)

		if err != nil {
			log.Println(err)

			reply.Type = events.AccommodationNotUpdate
		} else {
			Accommodation.AverageRating = average
			reply.Type = events.AccommodationUpdate
		}

	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
