package api

import (
	"github.com/tamararankovic/microservices_demo/common/notification"
	events "github.com/tamararankovic/microservices_demo/common/saga/delete_rating"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"log"
	"rating_service/application"
	"rating_service/domain"
)

type DeleteRatingCommandHandler struct {
	ratingService         *application.RatingService
	replyPublisher        saga.Publisher
	commandSubscriber     saga.Subscriber
	notificationPublisher saga.Publisher
}

func NewDeleteRatingCommandHandler(ratingService *application.RatingService, publisher saga.Publisher, subscriber saga.Subscriber, notificationPublisher saga.Publisher) (*DeleteRatingCommandHandler, error) {
	o := &DeleteRatingCommandHandler{
		ratingService:         ratingService,
		replyPublisher:        publisher,
		commandSubscriber:     subscriber,
		notificationPublisher: notificationPublisher,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteRatingCommandHandler) handle(command *events.DeleteRatingCommand) {
	reply := events.DeleteRatingReply{Rating: command.Rating}
	switch command.Type {
	case events.StartedDeletionRating:
		err := handler.ratingService.DeleteRating(command.Rating.ID)
		if err != nil {

			reply.Type = events.CantFindRating

		} else {
			reply.Type = events.DeletionStarted
		}
	case events.RollbackRating:
		oldValue := command.Rating.OldValue
		r := domain.Rating{
			Id:           oldValue.ID,
			UserID:       oldValue.UserID,
			TargetId:     oldValue.TargetID,
			RatingValue:  int32(oldValue.Value),
			TargetType:   int(oldValue.TargetType),
			LastModified: oldValue.LastModified,
		}
		err := handler.ratingService.Create(&r)
		if err != nil {
			log.Println(err)
		}

		reply.Type = events.RatingRollBack
	case events.SendNotification:
		log.Println("RATING DELETED SUCCESSFULLY")
		handler.notificationPublisher.Publish(&notification.Message{
			Title:      "Rating Created Successfully",
			Content:    "Rating Created Successfully",
			Type:       notification.AccommodationRated,
			NotifierId: (*(*command).Rating.OldValue).UserID,
		})
		reply.Type = events.RatingDeletionDone
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
