package api

import (
	"fmt"
	"github.com/tamararankovic/microservices_demo/common/notification"
	events "github.com/tamararankovic/microservices_demo/common/saga/leave_rating"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"log"
	"rating_service/application"
	"rating_service/domain"
)

type LeaveRatingCommandHandler struct {
	ratingService         *application.RatingService
	replyPublisher        saga.Publisher
	commandSubscriber     saga.Subscriber
	notificationPublisher saga.Publisher
}

func NewLeaveRatingCommandHandler(ratingService *application.RatingService, publisher saga.Publisher, subscriber saga.Subscriber, notificationPublisher saga.Publisher) (*LeaveRatingCommandHandler, error) {
	o := &LeaveRatingCommandHandler{
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

func (handler *LeaveRatingCommandHandler) handle(command *events.LeaveRatingCommand) {

	reply := events.LeaveRatingReply{Rating: command.Rating}
	fmt.Println("4")
	fmt.Println(command.Type)
	switch command.Type {
	case events.CreateRating:
		fmt.Println("5")
		oldValue := command.Rating.OldValue
		var err error
		r := domain.Rating{
			Id:           command.Rating.ID,
			UserID:       command.Rating.UserID,
			TargetId:     command.Rating.TargetID,
			RatingValue:  int32(command.Rating.Value),
			TargetType:   int(command.Rating.TargetType),
			LastModified: command.Rating.LastModified,
		}
		fmt.Println("3")
		println(oldValue)
		fmt.Println(r)
		log.Println(r)

		if oldValue == nil {

			reply.Rating.ID = r.Id
			err = handler.ratingService.Create(&r)
		} else {
			r.Id = oldValue.ID
			reply.Rating.ID = r.Id
			err = (*handler.ratingService).Update(&r)
		}
		if err != nil {
			log.Println(err)
			reply.Type = events.CreationFailed
		} else {
			reply.Type = events.RatingCreated
		}
	case events.RollBackRating:
		oldValue := command.Rating.OldValue
		if oldValue == nil {
			handler.ratingService.Delete(&domain.Rating{Id: command.Rating.ID})
		} else {
			handler.ratingService.Update(&domain.Rating{
				Id:           oldValue.ID,
				TargetId:     oldValue.TargetID,
				UserID:       oldValue.UserID,
				TargetType:   int(oldValue.TargetType),
				RatingValue:  int32(oldValue.Value),
				LastModified: command.Rating.LastModified,
			})
		}
		log.Println("RATING ROLLED BACK")
		reply.Type = events.RatingRollBack
	case events.SendNotification:
		fmt.Println("send notifcation")
		log.Println("notification sent")
		handler.notificationPublisher.Publish(&notification.Message{
			Title:      "Rating Created Successfully",
			Content:    "Rating Created Successfully",
			Type:       notification.AccommodationRated,
			NotifierId: command.Rating.UserID,
		})
		fmt.Println("mmmmm")
		reply.Type = events.NotificationSent
		return
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
