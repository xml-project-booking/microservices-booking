package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"term_service/domain"
	"time"
)

type TermService struct {
	store domain.TermStore
}

func NewTermService(store domain.TermStore) *TermService {
	return &TermService{
		store: store,
	}
}

func (service *TermService) Get(id primitive.ObjectID) (*domain.Term, error) {
	return service.store.Get(id)
}

func (service *TermService) GetAll() ([]*domain.Term, error) {
	return service.store.GetAll()
}

func (service *TermService) Create(term *domain.Term) error {
	term.Id = primitive.NewObjectID()
	return service.store.Insert(term)
}

func (service *TermService) Cancel(term *domain.Term) error {
	return service.store.Update(term)
}
func (service *TermService) Update(term *domain.Term) error {
	return service.store.Update(term)
}

func (service *TermService) Delete(term *domain.Term) error {
	return service.store.Delete(term)
}

func (service *TermService) DeleteReservationsInDateRange(accommodationId primitive.ObjectID, startDate time.Time, endDate time.Time) bool {

	allTerms, _ := service.store.GetAll()
	var filteredTerms []*domain.Term

	for _, term := range allTerms {
		if term.AccommodationID == accommodationId {
			filteredTerms = append(filteredTerms, term)
		}
	}

	for _, res := range filteredTerms {
		var isTaken = service.CheckIfOverLaps(res.Date, startDate, endDate)

		if isTaken == true {
			if res.UserID.String() != "" {
				service.Delete(res)
			} else {
				return false
			}

		}

	}
	return true

}

func (service *TermService) CheckForReservationInDateRange(accommodationId primitive.ObjectID, startDate time.Time, endDate time.Time) bool {

	allTerms, _ := service.store.GetAll()
	var filteredTerms []*domain.Term

	for _, term := range allTerms {
		if term.AccommodationID == accommodationId {
			filteredTerms = append(filteredTerms, term)
		}
	}

	for _, res := range filteredTerms {
		var isTaken = service.CheckIfOverLaps(res.Date, startDate, endDate)
		if isTaken == true {
			return isTaken
		}

	}
	return false

}

func (service *TermService) CheckIfOverLaps(date time.Time, start time.Time, end time.Time) bool {
	if start.Before(date) && (end.After(date)) {
		return true
	}
	return false
}

func (service *TermService) GetByAccommodationId(id primitive.ObjectID) ([]*domain.Term, error) {
	allTerms, _ := service.store.GetAll()
	var filteredTerms []*domain.Term

	for _, term := range allTerms {
		if term.AccommodationID == id {
			filteredTerms = append(filteredTerms, term)
		}
	}

	return filteredTerms, nil
}
