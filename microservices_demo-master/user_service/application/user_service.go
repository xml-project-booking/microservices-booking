package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user_service/domain"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetProminentHosts() ([]*domain.User, error) {
	return service.store.GetProminentHosts()
}

func (service *UserService) UpdateCancellationNumber(user *domain.User) error {
	return service.store.UpdateCancellationNumber(user)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Create(user *domain.User) error {
	return service.store.Insert(user)
}

func (service *UserService) Cancel(user *domain.User) error {
	return service.store.UpdateStatus(user)
}
func (service *UserService) UpdateUser(user *domain.User) error {
	return service.store.UpdateUser(user)
}
func (service *UserService) DeleteUser(id string) error {
	return service.store.DeleteUser(id)
}
