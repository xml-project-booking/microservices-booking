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
	GuestId            string
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

type UserInfo struct {
	Id   string
	Role string
}
type TermDTO struct {
	Id              string
	AccommodationID string // ID smeštaja za koji je termin vezan
	UserID          string // ID korisnika koji je rezervisao smeštaj
	PriceType       string
	Value           string
	StartDate       string
	EndDate         string
}
