package application

import (
	"accommodation_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingService struct {
	store domain.RatingStore
}

func NewRatingService(store domain.RatingStore) *RatingService {
	return &RatingService{
		store: store,
	}
}

func (service *RatingService) GetAccommodationAverage(accommodationId primitive.ObjectID) (float64, error) {
	accommodationRatings, err := service.store.GetTargetRatings(accommodationId, 0)
	if err != nil {
		return 0, err
	}
	var sum int32
	for _, rating := range accommodationRatings {
		if rating.RatingValue == 0 || rating.RatingValue == -1 || rating.RatingValue == 1 { // Check for special cases if needed
			return 0.0, errors.New("invalid input: division by zero or special value")
		}
		sum += rating.RatingValue
	}

	return float64(sum / int32(len(accommodationRatings))), nil

}

func (service *RatingService) GetHostAverage(hostId primitive.ObjectID) (float64, error) {
	hostRatings, err := service.store.GetTargetRatings(hostId, 1)
	if err != nil {
		return 0, err
	}
	var sum int32
	for _, rating := range hostRatings {
		if rating.RatingValue == 0 || rating.RatingValue == -1 || rating.RatingValue == 1 { // Check for special cases if needed
			return 0.0, errors.New("invalid input: division by zero or special value")
		}
		sum += rating.RatingValue
	}

	return float64(sum / int32(len(hostRatings))), nil
}

func (service *RatingService) Get(id primitive.ObjectID) (*domain.Rating, error) {
	return service.store.Get(id)
}

func (service *RatingService) GetAll() ([]*domain.Rating, error) {
	return service.store.GetAll()
}

func (service *RatingService) Create(term *domain.Rating) error {
	term.Id = primitive.NewObjectID()
	return service.store.Insert(term)
}

func (service *RatingService) Cancel(term *domain.Rating) error {
	return service.store.Update(term)
}
func (service *RatingService) Update(term *domain.Rating) error {
	return service.store.Update(term)
}

func (service *RatingService) Delete(term *domain.Rating) error {
	return service.store.Delete(term)
}
