package application

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user_service/domain"
)

var ErrorUsernameTaken = errors.New("Username is already taken")
var ErrorUserNotFound = errors.New("User not found")
var ErrorInvalidPassword = errors.New("Password is invalid")

type AuthentificationService struct {
	store domain.UserStore
}

func NewAuthentificationService(store domain.UserStore) *AuthentificationService {
	return &AuthentificationService{
		store: store,
	}
}

func (service *AuthentificationService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}
func (service *AuthentificationService) ExistsUsername(username string) (bool, error) {
	return service.store.CheckIfUsernameExists(username)
}
func (service *AuthentificationService) ExistsEmail(email string) (bool, error) {
	return service.store.CheckIfEmailExists(email)
}

func (service *AuthentificationService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}
func (service *AuthentificationService) RegisterUser(user *domain.User) (jwtToken string, err error) {
	exists, err := service.store.CheckIfUsernameExists(user.Username)

	if err != nil {
		return
	}

	if exists {
		return jwtToken, ErrorUsernameTaken
	}

	user.Password, err = HashPassword(user.Password)

	if err != nil {
		return
	}

	err = service.store.Insert(user)

	if err != nil {
		return
	}
	stringRole := "HOST"
	if user.Role == 1 {
		stringRole = "GUEST"
	}

	return GenerateJWT(user.Username, user.Id.String(), stringRole)
}
func (service *AuthentificationService) Login(username string, password string) (jwtToken string, role string, id string, err error) {
	user, err := service.store.GetUserByUsername(username)

	if err != nil {
		return jwtToken, role, id, ErrorUserNotFound
	}
	if user == nil {
		return jwtToken, role, id, ErrorUserNotFound
	}

	if !CheckPasswordHash(password, user.Password) {
		return jwtToken, role, id, ErrorInvalidPassword
	}
	stringRole := "HOST"
	if user.Role == 1 {
		stringRole = "GUEST"
	}

	jwtToken, err = GenerateJWT(user.Username, user.Id.String(), stringRole)

	if err != nil {
		return jwtToken, role, id, err
	}

	return jwtToken, stringRole, user.Id.String(), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GenerateJWT(username, user_id, role string) (string, error) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString, nil
}
