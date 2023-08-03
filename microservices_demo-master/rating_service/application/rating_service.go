package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_service/domain"
)

type RatingService struct {
	store domain.RatingStore
}

func NewRatingService(store domain.RatingStore) *RatingService {
	return &RatingService{
		store: store,
	}
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
