package persistence

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"

	"user_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "users"
	COLLECTION = "users"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

var ErrorUsernameTaken = errors.New("Username is already taken")

func (store *UserMongoDBStore) UpdateStatus(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
func (store *UserMongoDBStore) UpdateCancellationNumber(user *domain.User) error {
	number := user.CancellationNumber + 1
	result, err := store.users.UpdateOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		bson.D{
			{"$set", bson.D{{"cancellation_number", number}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("one document should've been updated")
	}
	return nil

}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Insert(User *domain.User) error {
	result, err := store.users.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}
	User.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) DeleteAll() {
	store.users.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserMongoDBStore) filterOne(filter interface{}) (User *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func (store *UserMongoDBStore) GetUserByUsername(username string) (User *domain.User, err error) {
	filter := bson.D{{"username", username}}
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func (store *UserMongoDBStore) CheckIfUsernameExists(username string) (bool, error) {
	filter := bson.D{{"username", username}}
	var _, err = store.filterOne(filter)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return true, err
	//filter := bson.D{{"_id", id}, {"deleted", false}}
	//result, err := store.filterOne(filter)
	//
	//if err != nil {
	//	if err == mongo.ErrNoDocuments {
	//		return nil, persistance.ErrorUserNotFound
	//	}
	//	return nil, err
	//}
	//
	//return result, nil
}
func (store *UserMongoDBStore) CheckIfEmailExists(email string) (bool, error) {
	filter := bson.D{{"email", email}}
	var _, err = store.filterOne(filter)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return true, err
}
func (store *UserMongoDBStore) UpdateUser(user *domain.User) error {
	filter := bson.M{"_id": user.Id}
	result, err2 := store.filterOne(filter)
	if err2 != nil {
		return err2
	}
	user.Role = result.Role
	if user.Password == "" {
		user.Password = result.Password
	} else {
		user.Password, _ = HashPassword(user.Password)
	}
	if user.Username != result.Username {
		exists, err := store.CheckIfUsernameExists(user.Username)
		if err != nil {
			return err
		}

		if exists {
			return ErrorUsernameTaken
		}
	}
	_, err := store.users.ReplaceOne(context.TODO(), bson.M{"_id": user.Id}, user)

	if err != nil {
		return err
	}

	return nil
}
func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var User domain.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
