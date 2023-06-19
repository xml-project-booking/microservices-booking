package persistence

import (
	"context"
	"term_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "terms"
	COLLECTION = "terms"
)

type TermMongoDBStore struct {
	terms *mongo.Collection
}

func (store *TermMongoDBStore) GetByAccommodationId(id primitive.ObjectID) ([]*domain.Term, error) {
	filter := bson.M{"accommodationId": id}
	return store.filter(filter)
}

func (store *TermMongoDBStore) Delete(term *domain.Term) error {

	_, err := store.terms.DeleteOne(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *TermMongoDBStore) Update(term *domain.Term) error {

	_, err := store.terms.ReplaceOne(context.TODO(), bson.M{"_id": term.Id}, term)

	if err != nil {
		return err
	}

	return nil
}

func NewTermMongoDBStore(client *mongo.Client) domain.TermStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &TermMongoDBStore{
		terms: users,
	}
}

func (store *TermMongoDBStore) Get(id primitive.ObjectID) (*domain.Term, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *TermMongoDBStore) GetAll() ([]*domain.Term, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *TermMongoDBStore) Insert(User *domain.Term) error {
	result, err := store.terms.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}
	User.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *TermMongoDBStore) DeleteAll() {
	store.terms.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *TermMongoDBStore) filter(filter interface{}) ([]*domain.Term, error) {
	cursor, err := store.terms.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *TermMongoDBStore) filterOne(filter interface{}) (User *domain.Term, err error) {
	result := store.terms.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}

func decode(cursor *mongo.Cursor) (terms []*domain.Term, err error) {
	for cursor.Next(context.TODO()) {
		var Term domain.Term
		err = cursor.Decode(&Term)
		if err != nil {
			return
		}
		terms = append(terms, &Term)
	}
	err = cursor.Err()
	return
}
