package persistence

import (
	"accommodation_service/domain"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accommodationDB"
	COLLECTION = "accommodation"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func (store *AccommodationMongoDBStore) FilterAccommodationsByAmenities(amenities []bool) ([]*domain.Accommodation, error) {
	filter := bson.M{"wifi": amenities[0], "kitchen": amenities[1], "air_conditioning": amenities[2], "free_parking": amenities[3]}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) UpdateStatus(user *domain.Accommodation) error {
	//TODO implement me
	panic("implement me")
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain.AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) Get(id primitive.ObjectID) (*domain.Accommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) GetAll() ([]*domain.Accommodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) UpdateReservationConfirmationType(accommodation *domain.Accommodation) error {
	result, err := store.accommodations.UpdateOne(
		context.TODO(),
		bson.M{"_id": accommodation.Id},
		bson.D{
			{"$set", bson.D{{"reservation_confirmation", accommodation.ReservationConfirmation}}},
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

func (store *AccommodationMongoDBStore) Insert(Accommodation *domain.Accommodation) error {
	Accommodation.Id = primitive.NewObjectID()
	result, err := store.accommodations.InsertOne(context.TODO(), &Accommodation)
	fmt.Println(result)
	if err != nil {
		return err
	}
	Accommodation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AccommodationMongoDBStore) DeleteAll() {
	store.accommodations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*domain.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (User *domain.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&User)
	return
}
func (store *AccommodationMongoDBStore) DeleteAccommodation(id primitive.ObjectID) error {
	_, err := store.accommodations.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return err
	}

	return nil
}

func decode(cursor *mongo.Cursor) (accommodations []*domain.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var Accommodation domain.Accommodation
		err = cursor.Decode(&Accommodation)
		if err != nil {
			return
		}
		accommodations = append(accommodations, &Accommodation)
	}
	err = cursor.Err()
	return
}
