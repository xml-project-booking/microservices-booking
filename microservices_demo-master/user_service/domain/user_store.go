package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) error
	DeleteAll()
	UpdateStatus(user *User) error
	CheckIfUsernameExists(username string) (bool, error)
	GetUserByUsername(username string) (*User, error)
	CheckIfEmailExists(email string) (bool, error)
	UpdateUser(user *User) error
	UpdateCancellationNumber(user *User) error
}
