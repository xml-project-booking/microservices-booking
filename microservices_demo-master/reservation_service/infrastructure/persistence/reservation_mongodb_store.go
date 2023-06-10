package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"resevation/domain"
)

const (
	DATABASE   = "reservations"
	COLLECTION = "reservations"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func (store *ReservationMongoDBStore) UpdateStatus(user *domain.Reservation) error {
	//TODO implement me
	panic("implement me")
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (store *ReservationMongoDBStore) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ReservationMongoDBStore) GetAllReservationRequests() ([]*domain.Reservation, error) {
	filter := bson.M{"confirmation": false}
	return store.filter(filter)
}
func (store *ReservationMongoDBStore) GetAllReservation() ([]*domain.Reservation, error) {
	filter := bson.M{"confirmation": true}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) GetAll() ([]*domain.Reservation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) Insert(Reservation *domain.Reservation) error {
	fmt.Println("dosao si dovdee majnmuneeeeee")
	fmt.Println(Reservation)
	Reservation.Id = primitive.NewObjectID()
	result, err := store.reservations.InsertOne(context.TODO(), &Reservation)
	fmt.Println(result)
	if err != nil {
		return err
	}
	Reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationMongoDBStore) DeleteAll() {
	store.reservations.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *ReservationMongoDBStore) filter(filter interface{}) ([]*domain.Reservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *ReservationMongoDBStore) filterOne(filter interface{}) (Reservation *domain.Reservation, err error) {
	result := store.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&Reservation)
	return
}

func decode(cursor *mongo.Cursor) (reservations []*domain.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var Reservation domain.Reservation
		err = cursor.Decode(&Reservation)
		if err != nil {
			return
		}
		reservations = append(reservations, &Reservation)
	}
	err = cursor.Err()
	return
}
