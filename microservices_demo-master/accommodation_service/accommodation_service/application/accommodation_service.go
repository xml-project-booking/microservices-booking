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

func (service *AccommodationService) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	return service.store.Get(id)
}

func (service *AccommodationService) GetAll() ([]*domain.Accommodation, error) {
	return service.store.GetAll()
}

func (service *AccommodationService) Create(user *domain.Accommodation) error {
	return service.store.Insert(user)
}

func (service *AccommodationService) Cancel(user *domain.Accommodation) error {
	return service.store.UpdateStatus(user)
}
func (service *AccommodationService) GetAllAccommodationsByHost(id string) []domain.Accommodation {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	filteredAccommodations := []domain.Accommodation{}
	accomodations, _ := service.store.GetAll()
	for _, accomodation := range accomodations {
		if accomodation.HostId == objectId {
			filteredAccommodations = append(filteredAccommodations, *accomodation)
		}
	}
	return filteredAccommodations
}

func (service *AccommodationService) DeleteAllAccommodationsByHost(id string) error {
	for _, accommodation := range service.GetAllAccommodationsByHost(id) {
		err := service.store.DeleteAccommodation(accommodation.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
