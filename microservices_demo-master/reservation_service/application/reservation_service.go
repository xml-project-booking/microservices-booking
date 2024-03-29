package application

import (
	"errors"
	"fmt"
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
func (service *ReservationService) GetAllReservationsByAccommodation(accommodationId primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetAllReservationByAccommodation(accommodationId)
}
func (service *ReservationService) GetAllConfirmedReservationsByAccommodation(accommodationId primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetAllConfirmedReservationByAccommodation(accommodationId)
}
func (service *ReservationService) GetAllReservationsByGuestId(guestId primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetAllReservationByGuest(guestId)
}
func (service *ReservationService) GetAllReservationsByGuestIdPending(guestId primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetAllReservationByGuestPending(guestId)
}
func (service *ReservationService) GetAllReservation() ([]*domain.Reservation, error) {
	return service.store.GetAllReservation()
}
func (service *ReservationService) CheckTotalReservationDuration(hostId primitive.ObjectID) bool {
	reservations, _ := service.store.GetReservationsByHost(hostId)
	var duration = 0.0
	for _, Reservation := range reservations {
		difference := (Reservation.EndDate.Sub(Reservation.StartDate).Hours()) / 24
		duration = duration + difference
	}
	if duration > 50 {
		return true
	} else {
		return false
	}
}
func (service *ReservationService) CheckReservationNumberForHost(hostId primitive.ObjectID) bool {
	reservations, _ := service.store.GetReservationsByHost(hostId)
	var num = len(reservations)
	if num >= 5 {
		return true
	} else {
		return false
	}
}
func (service *ReservationService) CheckCancellationRate(hostId primitive.ObjectID) bool {
	hostCanceledReservations, _ := service.store.GetReservationCancelByHost(hostId)
	hostReservations, _ := service.store.GetReservationsByHost(hostId)
	if len(hostReservations) == 0 {
		return true
	}
	cancellationRate := (len(hostCanceledReservations) / len(hostReservations)) * 100

	if cancellationRate < 5 {
		return true
	} else {
		return false
	}

}
func (service *ReservationService) GetAll() ([]*domain.Reservation, error) {
	return service.store.GetAll()
}
func (service *ReservationService) GetAllConfirmedReservations() ([]*domain.Reservation, error) {
	return service.store.GetAllReservation()
}

func (service *ReservationService) Create(user *domain.Reservation) error {
	return service.store.Insert(user)
}

func (service *ReservationService) GetAllGuestReservations(guestId primitive.ObjectID) ([]*domain.Reservation, error) {
	return service.store.GetAllGuestReservation(guestId)
}
func (service *ReservationService) UpdateReservationStatusForCanceled(reservation *domain.Reservation) error {
	return service.store.UpdateStatusForCanceled(reservation)

}
func (service *ReservationService) UpdateReservationStatusForConfirmed(reservation *domain.Reservation) error {
	return service.store.UpdateStatusForConfirmed(reservation)

}
func (service *ReservationService) Cancel(user *domain.Reservation) error {
	return service.store.UpdateStatusForCanceled(user)
}
func (service *ReservationService) Delete(reservationId primitive.ObjectID) bool {
	return service.store.DeleteReservationById(reservationId)
}
func (service *ReservationService) CreateReservationRequest(reservation *domain.Reservation) error {
	err := service.store.Insert(reservation)
	if err != nil {
		return err
	}
	return nil
}
func (service *ReservationService) CancelAllReservationRequestOverlaping(reservations []*domain.Reservation) {
	for _, res := range reservations {
		service.UpdateReservationStatusForCanceled(res)
	}

}
func (service *ReservationService) GetAllReservationsSearch(startDate time.Time, endDate time.Time) []*domain.Reservation {
	allReservations, _ := service.store.GetAllReservation()
	fmt.Println("ovo su svi zahtjevi za rezeervaciju")
	fmt.Println(allReservations)
	var filteredReservationsOverlaps []*domain.Reservation
	for _, res := range allReservations {
		fmt.Println(res)
		var overLaps = service.CheckIfOverLaps(res.StartDate, res.EndDate, startDate, endDate)
		if overLaps {
			fmt.Println("usao u fju za prekpalanje")
			filteredReservationsOverlaps = append(filteredReservationsOverlaps, res)
		}

	}
	return filteredReservationsOverlaps
}
func (service *ReservationService) GetAllReservationRequestWhichOverlapsWithConfirmed(startDate time.Time, endDate time.Time, accommodationId primitive.ObjectID) []*domain.Reservation {
	allReservations, _ := service.store.GetAllReservationByAccommodation(accommodationId)
	fmt.Println("ovo su svi zahtjevi za rezeervaciju")
	fmt.Println(allReservations)
	var filteredReservationsOverlaps []*domain.Reservation
	for _, res := range allReservations {
		fmt.Println(res)
		var overLaps = service.CheckIfOverLaps(res.StartDate, res.EndDate, startDate, endDate)
		if overLaps {
			fmt.Println("usao u fju za prekpalanje")
			filteredReservationsOverlaps = append(filteredReservationsOverlaps, res)
		}

	}
	return filteredReservationsOverlaps
}
func (service *ReservationService) CheckForReservationInDateRange(accommodationId primitive.ObjectID, startDate time.Time, endDate time.Time) bool {

	allReservations, _ := service.store.GetAllConfirmedReservationByAccommodation(accommodationId)
	for _, res := range allReservations {
		var isTaken = service.CheckIfOverLaps(res.StartDate, res.EndDate, startDate, endDate)
		if isTaken == true {
			return isTaken
		}

	}
	return false

}
func (service *ReservationService) CheckForReservationInDateRangeAndGetUserId(accommodationId primitive.ObjectID, startDate time.Time, endDate time.Time) string {
	fmt.Println("usaooooooooooo ovdee")
	allReservations, _ := service.store.GetAllConfirmedReservationByAccommodation(accommodationId)
	var termCondition string
	fmt.Println("lista filterovana po acc id")
	fmt.Println(allReservations)

	for _, res := range allReservations {
		var isTaken = service.CheckIfOverLaps(res.StartDate, res.EndDate, startDate, endDate)
		if isTaken {
			fmt.Println("uslov ispunjen")
			termCondition = res.Id.Hex()
			return termCondition
		}

	}

	termCondition = "greska"

	fmt.Println(termCondition)
	return termCondition

}

func (service *ReservationService) CancelReservation(reservation *domain.Reservation) string {

	if service.CheckIfLessThan24Hours(reservation.Id) {
		return errors.New("manje od 24h").Error()
	}

	var isDeleted = service.store.UpdateStatusForCanceledUser(reservation)
	if isDeleted != nil {
		return isDeleted.Error()
	}
	return "proslo"

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

func (service *ReservationService) CheckIfLessThan24Hours(reservationId primitive.ObjectID) bool {
	reservation, _ := service.store.Get(reservationId)
	startTime := time.Now()
	fmt.Println("ovo je rez")
	fmt.Println(reservation)

	startReservationTime := reservation.StartDate
	difference := (startReservationTime).Sub(startTime).Hours()
	if difference < 24 {
		return true
	}
	return false

}

func (service *ReservationService) CheckIfNumberOfGuestIsValid(minNumber int64, maxNumber int64, guestNumber int64) bool {
	return guestNumber >= minNumber && guestNumber <= maxNumber
}

func (service *ReservationService) CheckGuestCanLeaveRating(accommodationId, guestId primitive.ObjectID) bool {
	var reservations, _ = service.store.GetGuestAccommodationReservation(accommodationId, guestId)
	var pastReservations []*domain.Reservation
	for _, Reservation := range reservations {
		if Reservation.EndDate.Before(time.Now()) {
			pastReservations = append(pastReservations, Reservation)
		}
	}
	var num = len(pastReservations)
	if num > 0 {
		return true
	}
	return false
}
func (service *ReservationService) CheckGuestCanLeaveRatingForHost(hostId, guestId primitive.ObjectID) bool {
	var reservations, _ = service.store.GetGuestAccommodationHostReservation(hostId, guestId)
	var pastReservations []*domain.Reservation
	for _, Reservation := range reservations {
		if Reservation.EndDate.Before(time.Now()) {
			pastReservations = append(pastReservations, Reservation)
		}
	}
	var num = len(pastReservations)
	if num > 0 {
		return true
	}
	return false
}
