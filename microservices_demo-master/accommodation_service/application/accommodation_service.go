package application

import (
	"accommodation_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}
func (service *AccommodationService) SearchAccommodationsByLocation(accommodations []*domain.Accommodation, location string) []*domain.Accommodation {
	var filterAccommodations []*domain.Accommodation
	for _, Accommodation := range accommodations {
		result := strings.Contains(strings.ToLower(Accommodation.City), strings.ToLower(location))
		resultOne := strings.Contains(strings.ToLower(Accommodation.Country), strings.ToLower(location))
		if result || resultOne {
			filterAccommodations = append(filterAccommodations, Accommodation)
		}
	}
	return filterAccommodations
}
func (service *AccommodationService) CheckAccommodationForAmenities(amenities []bool, accommodations []*domain.Accommodation) ([]*domain.Accommodation, error) {
	var filteredAccommodations []*domain.Accommodation
	for _, Accommodation := range accommodations {
		if Accommodation.Wifi == amenities[0] && Accommodation.Kitchen == amenities[1] && Accommodation.AirConditioning == amenities[2] && Accommodation.FreeParking == amenities[3] {
			filteredAccommodations = append(filteredAccommodations, Accommodation)
		}
	}
	return filteredAccommodations, nil

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
