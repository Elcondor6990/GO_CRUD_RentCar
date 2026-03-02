package repository

import (
	"Go_Rent_Car/src/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type reservationCarRepo struct {
	collection *mongo.Collection
}

type ReservationCarRepository interface{
	Create(ctx context.Context, reserve *models.ReservationCar)(*models.ReservationCar, error)
	FindByID(ctx context.Context, id string)(*models.ReservationCar, error)
	Delete(ctx context.Context, id string) error
}


func NewReservationCarRepository(db *mongo.Database) ReservationCarRepository{
	return &reservationCarRepo{collection: db.Collection("reservations_car")}
}


func (r *reservationCarRepo) Create(ctx context.Context, reserve *models.ReservationCar)(*models.ReservationCar, error){
	create, err := r.collection.InsertOne(ctx, reserve)

	if err != nil{
		return nil, err
	}

	reserve.ID = create.InsertedID.(primitive.ObjectID)
	return reserve, nil 
}

func (r *reservationCarRepo) FindByID(ctx context.Context, id string)(*models.ReservationCar, error){
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid ID")
	}

	var reserveCar models.ReservationCar

	err = r.collection.FindOne(ctx, bson.M{"_id": objectid}).Decode(&reserveCar)
	if err != nil{
		return nil, err
	}

	return &reserveCar, nil
}


func (r *reservationCarRepo) Delete(ctx context.Context, id string) error{
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid ID")
	}
	
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectid})
	if err != nil{
		return err
	}

	if result.DeletedCount == 0{
		return errors.New("Reservation not found")
	}
	return nil
}