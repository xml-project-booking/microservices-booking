package delete_rating

type Color struct {
	Code string
}

type Product struct {
	Id    string
	Color Color
}

type OrderItem struct {
	Product  Product
	Quantity uint16
}

type RatingDetails struct {
	Id string
}

type DeleteRatingCommandType int8

const (
	StartedDeletionRating DeleteRatingCommandType = iota
	UpdateAccommodation
	RollbackRating
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
	RatingRollBack
	RatingDeletionDone
	UnknownReply
)

type DeleteRatingReply struct {
	Rating RatingDetails
	Type   DeleteRatingReplyType
}
