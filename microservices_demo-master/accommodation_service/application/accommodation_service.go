package application

import (
	"accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) UpdateReservationConfirmationType(accommodation *domain.Accommodation) error {
	return service.store.UpdateReservationConfirmationType(accommodation)
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(accommodation *domain.Accommodation) error {
	return service.store.Insert(accommodation)

}

func (service *AccommodationService) Cancel(user *domain.Accommodation) error {
	return service.store.UpdateStatus(user)
}
