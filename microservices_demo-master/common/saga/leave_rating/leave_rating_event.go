package leave_rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RatingDetails struct {
	ID           primitive.ObjectID
	TargetID     primitive.ObjectID
	TargetType   uint32
	UserID       primitive.ObjectID
	Value        uint32
	LastModified time.Time
	OldValue     *RatingDetails
}

type LeaveRatingCommandType int8

const (
	StartedCreatingRating LeaveRatingCommandType = iota
	CreateRating
	UpdateAccommodation
	ApproveRating
	RollBackRating
	SendNotification

	CancelRating
	UnknownCommand
)

type LeaveRatingCommand struct {
	Rating RatingDetails
	Type   LeaveRatingCommandType
}

type LeaveRatingReplyType int8

const (
	CreationStarted LeaveRatingReplyType = iota
	RatingCreated
	CreationFailed
	CantGiveRating
	AccommodationUpdate
	AccommodationNotUpdate
	RatingRollBack
	NotificationSent
	UnknownReply
)

type LeaveRatingReply struct {
	Rating RatingDetails
	Type   LeaveRatingReplyType
}
