package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"term_service/domain"
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

func (service *TermService) Create(user *domain.Term) error {
	return service.store.Insert(user)
}

func (service *TermService) Cancel(user *domain.Term) error {
	return service.store.UpdateStatus(user)
}
