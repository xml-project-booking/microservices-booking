package application

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/term_service"
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

func (service *TermService) GetAvailableAccommodationsInPeriod(startDate string, endDate string) ([]*pb.AccommodationWithPriceResponse, error) {
	allTerms, _ := service.store.GetAll()
	var accommodationWithPriceResponse []*pb.AccommodationWithPriceResponse
	layout := "2006-01-02T15:04:05.000Z"
	strtDate, _ := time.Parse(layout, startDate)
	eDate, _ := time.Parse(layout, endDate)

	var cantBeAddedIds []primitive.ObjectID

	//set hours, seconds... to 0
	strtDate = time.Date(strtDate.Year(), strtDate.Month(), strtDate.Day(), 0, 0, 0, 0, strtDate.Location())
	eDate = time.Date(eDate.Year(), eDate.Month(), eDate.Day(), 0, 0, 0, 0, eDate.Location())

	for _, term := range allTerms {

		termDateInPeriod := service.CheckIfOverLaps(term.Date, strtDate, eDate)
		if termDateInPeriod == false {
			continue
		}
		
		//If cantBeAddedIds contains term.AccommodationID
		var cantBeAdded = false
		for _, id := range cantBeAddedIds {
			cantBeAdded = (id == term.AccommodationID)
			if cantBeAdded == true {
				break
			}
		}
		if cantBeAdded == true {
			continue
		}

		//Is term reserved by user
		if term.UserID != primitive.NilObjectID {
			//Check if cantBeAddedIds contains term.AccommodationID
			var isAlreadyInSlice = false
			for _, id := range cantBeAddedIds {
				if id == term.AccommodationID {
					isAlreadyInSlice = true
				}
			}

			if isAlreadyInSlice == false {
				cantBeAddedIds = append(cantBeAddedIds, term.AccommodationID)
			}
			continue
		}

		if termDateInPeriod == true {
			var isAlreadyInSlice = false

			for _, id := range accommodationWithPriceResponse {
				if id.AccommodationId == term.AccommodationID.Hex() {
					isAlreadyInSlice = true
				}
			}

			if isAlreadyInSlice == false {

				var newAccommodationWithPriceResponse = &pb.AccommodationWithPriceResponse{
					AccommodationId: term.AccommodationID.Hex(),
					Price:           term.Value,
				}
				accommodationWithPriceResponse = append(accommodationWithPriceResponse, newAccommodationWithPriceResponse)
			}

			if isAlreadyInSlice == true {
				for _, id := range accommodationWithPriceResponse {
					if id.AccommodationId == term.AccommodationID.Hex() {
						id.Price += term.Value
						break
					}
				}
			}

		}
	}

	return accommodationWithPriceResponse, nil
}
