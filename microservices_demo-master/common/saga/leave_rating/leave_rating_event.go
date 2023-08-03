package leave_rating

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
	Id              string
	GuestId         string
	AccommodationId string
	Rating          int32
}

type LeaveRatingCommandType int8

const (
	StartedCreatingRating LeaveRatingCommandType = iota
	CreateRating
	UpdateAccommodation
	ApproveRating
	RollBackRating
	GetOldValue
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
	AccommodationUpdate
	AccommodationNotUpdate
	RatingRollBack
	OldValueReturned
	RatingDone
	UnknownReply
)

type LeaveRatingReply struct {
	Rating RatingDetails
	Type   LeaveRatingReplyType
}
