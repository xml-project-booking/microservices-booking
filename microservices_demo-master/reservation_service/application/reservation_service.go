package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"resevation/domain"
)

type ReservationService struct {
	store domain.ReservationStore
}

func NewReservationService(store domain.ReservationStore) *ReservationService {
	return &ReservationService{
		store: store,
	}
}

func (service *ReservationService) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	return service.store.Get(id)
}

func (service *ReservationService) GetAll() ([]*domain.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) Create(user *domain.Reservation) error {
	return service.store.Insert(user)
}

func (service *ReservationService) Cancel(user *domain.Reservation) error {
	return service.store.UpdateStatus(user)
}
