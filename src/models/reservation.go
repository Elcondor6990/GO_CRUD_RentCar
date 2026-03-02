package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationCar struct {
	ID 				 	primitive.ObjectID  `bson:"_id,omitempty"`
	StartReservation 	primitive.DateTime 	`bson:"start_reserve"`
	EndReservation 	 	primitive.DateTime	`bson:"end_reserve"`
	NumberPlate			string				`bson:"number_plate"`
}