package persistence

import (
	"context"
	"rating_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "ratings"
	COLLECTION = "ratings"
)

type RatingMongoDBStore struct {
	ratings *mongo.Collection
}

func (store *RatingMongoDBStore) GetByAccommodationId(id primitive.ObjectID) ([]*domain.Rating, error) {
	filter := bson.M{"accommodationId": id}
	return store.filter(filter)
}

func (store *RatingMongoDBStore) Delete(rating *domain.Rating) error {

	_, err := store.ratings.DeleteOne(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *RatingMongoDBStore) Update(rating *domain.Rating) error {

	_, err := store.ratings.ReplaceOne(context.TODO(), bson.M{"_id": rating.Id}, rating)

	if err != nil {
		return err
	}

	return nil
}

func NewRatingMongoDBStore(client *mongo.Client) domain.RatingStore {
	ratings := client.Database(DATABASE).Collection(COLLECTION)
	return &RatingMongoDBStore{
		ratings: ratings,
	}
}

func (store *RatingMongoDBStore) Get(id primitive.ObjectID) (*domain.Rating, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *RatingMongoDBStore) GetAll() ([]*domain.Rating, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *RatingMongoDBStore) Insert(Rating *domain.Rating) error {
	result, err := store.ratings.InsertOne(context.TODO(), Rating)
	if err != nil {
		return err
	}
	Rating.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *RatingMongoDBStore) DeleteAll() {
	store.ratings.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *RatingMongoDBStore) filter(filter interface{}) ([]*domain.Rating, error) {
	cursor, err := store.ratings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *RatingMongoDBStore) filterOne(filter interface{}) (Rating *domain.Rating, err error) {
	result := store.ratings.FindOne(context.TODO(), filter)
	err = result.Decode(&Rating)
	return
}

func decode(cursor *mongo.Cursor) (ratings []*domain.Rating, err error) {
	for cursor.Next(context.TODO()) {
		var Rating domain.Rating
		err = cursor.Decode(&Rating)
		if err != nil {
			return
		}
		ratings = append(ratings, &Rating)
	}
	err = cursor.Err()
	return
}
