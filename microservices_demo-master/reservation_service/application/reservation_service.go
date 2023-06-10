package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"resevation/domain"
	"time"
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

func (service *ReservationService) GetAllReservationRequests() ([]*domain.Reservation, error) {
	return service.store.GetAllReservationRequests()
}

func (service *ReservationService) GetAllReservation() ([]*domain.Reservation, error) {
	return service.store.GetAllReservation()
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
func (service *ReservationService) CreateReservationRequest(reservation *domain.Reservation) error {
	err := service.store.Insert(reservation)
	if err != nil {
		return err
	}
	return nil
}
func (service *ReservationService) CheckForReservationInDateRange(accommodationId primitive.ObjectID, startDate time.Time, endDate time.Time) bool {

	allReservations, _ := service.store.GetAllReservation()
	var filteredReservations []*domain.Reservation

	for _, reservation := range allReservations {
		if reservation.AccommodationID == accommodationId {
			filteredReservations = append(filteredReservations, reservation)
		}
	}

	for _, res := range filteredReservations {
		var isTaken = service.CheckIfOverLaps(res.StartDate, res.EndDate, startDate, endDate)
		if isTaken == true {
			return isTaken
		}

	}
	return false

}

func (service *ReservationService) CheckIfOverLaps(start1 time.Time, end1 time.Time, start2 time.Time, end2 time.Time) bool {
	if (start1.Before(end2) || start1.Equal(end2)) && (end1.After(start2) || end1.Equal(start2)) {
		return true
	}

	if start1.Before(start2) && end1.After(end2) {
		return true
	}

	if start2.Before(start1) && end2.After(end1) {
		return true
	}

	return false
}
