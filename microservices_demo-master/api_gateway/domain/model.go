package domain

import (
	"time"
)

type Product struct {
	Id            string
	Name          string
	ClothingBrand string
	ColorCode     string
	ColorName     string
}

type OrderItem struct {
	Product  Product
	Quantity uint16
}

type OrderDetails struct {
	Id              string
	CreatedAt       time.Time
	Status          string
	ShippingAddress string
	ShippingStatus  string
	OrderItems      []OrderItem
}

type CancelReservation struct {
	UserId        string
	ReservationId string
}

type ReservationForHost struct {
	Id                 string
	StartDate          string
	EndDate            string
	GuestNumber        int64
	Confirmation       bool
	CancellationNumber int64
	GuestName          string
	GuestSurname       string
}
type ReservationDTO struct {
	Id              string
	AccommodationID string
	StartDate       string
	EndDate         string
	GuestNumber     string
	Confirmation    string
	GuestId         string
}
