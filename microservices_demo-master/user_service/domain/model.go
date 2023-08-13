package domain

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type User struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name               string             `bson:"name" json:"name"`
	Surname            string             `bson:"surname,omitempty" json:"surname"`
	Email              string             `bson:"email" json:"email"`
	Username           string             `bson:"username,omitempty" json:"username"`
	Password           string             `bson:"password" json:"password"`
	Address            string             `bson:"address,omitempty" json:"address"`
	Role               Role               `bson:"role" json:"role"`
	CancellationNumber int                `bson:"cancellation_number"`
	IsProminent        bool               `bson:"is_prominent"`
}

type Role int

const (
	Host = iota
	Guest
)

type Users []*User

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
	Role        string `json:"role"`
	Username    string `json:"username"`
	TokenString string `json:"token"`
}

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
