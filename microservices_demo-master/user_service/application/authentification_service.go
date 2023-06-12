package application

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user_service/domain"
)

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

func (service *AuthentificationService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}
func (service *AuthentificationService) RegisterUser(user *domain.User) (jwtToken string, err error) {
	exists, err := service.store.CheckIfUsernameExists(user.Username)

	if err != nil {
		return
	}

	if exists {
		return
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
func (service *AuthentificationService) Login(username string, password string) (jwtToken string, err error) {
	user, err := service.store.GetUserByUsername(username)

	if err != nil {
		return jwtToken, err
	}

	if !CheckPasswordHash(user.Password, password) {
		return jwtToken, err
	}
	stringRole := "HOST"
	if user.Role == 1 {
		stringRole = "GUEST"
	}

	jwtToken, err = GenerateJWT(user.Username, user.Id.String(), stringRole)

	if err != nil {
		return jwtToken, err
	}

	return jwtToken, nil
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
