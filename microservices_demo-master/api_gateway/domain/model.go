package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	HostId          string
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
type Accommodation struct {
	Id                      primitive.ObjectID `json:"id"`
	Name                    string             `json:"name"`
	ReservationConfirmation string             `json:"reservation_confirmation"`
	City                    string             `json:"city"`
	Country                 string             `json:"country"`
	Street                  string             `json:"street"`
	StreetNumber            string             `json:"street_number"`
	MinGuest                int                `json:"MinGuest"`
	MaxGuest                int                `json:"MaxGuest"`
	HostId                  primitive.ObjectID `json:"hostId"`
	Wifi                    bool               `json:"wifi"`
	Kitchen                 bool               `json:"kitchen"`
	AirConditioning         bool               `json:"airConditioning"`
	FreeParking             bool               `json:"freeParking"`
	AverageRating           float64            `json:"average_rating"`
	Price                   int64              `json:"price"`
	TotalPrice              int64              `json:"totalPrice"`
	Type                    string             `json:"type"`
}

type SearchDTO struct {
	StartDate   string
	EndDate     string
	Location    string
	GuestNumber int
}
type Mix struct {
	Amenities      []bool
	Accommodations []*Accommodation
}

type FilterParameters struct {
	Amenities      []bool
	Accommodations []*Accommodation
	MinPrice       int64
	MaxPrice       int64
	IsHost         bool
}
