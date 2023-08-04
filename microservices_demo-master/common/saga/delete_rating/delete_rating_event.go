package delete_rating

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

type DeleteRatingCommandType int8

const (
	StartedDeletionRating DeleteRatingCommandType = iota
	UpdateAccommodation
	UpdateHost
	RollbackRating
	CancelDeletionRating
	FinishDeletionRating
	UnknownCommand
)

type DeleteRatingCommand struct {
	Rating RatingDetails
	Type   DeleteRatingCommandType
}

type DeleteRatingReplyType int8

const (
	DeletionStarted DeleteRatingReplyType = iota
	DeletionFailed
	AccommodationUpdate
	AccommodationNotUpdate
	HostUpdate
	HostNotUpdate
	RatingRollBack
	RatingDeletionDone
	UnknownReply
)

type DeleteRatingReply struct {
	Rating RatingDetails
	Type   DeleteRatingReplyType
}
