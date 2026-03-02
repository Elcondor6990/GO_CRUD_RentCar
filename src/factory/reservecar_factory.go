package factory

import (
	"Go_Rent_Car/src/dto"
	"Go_Rent_Car/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModelToDto(reserve models.ReservationCar) dto.ReservationCarDto {
	return dto.ReservationCarDto{
		StartReservation: reserve.StartReservation.Time().String(),
		EndReservation:   reserve.EndReservation.Time().String(),
		NumberPlate:      reserve.NumberPlate,
	}
}

func DtoToModel(input dto.ReservationCarDto) models.ReservationCar {
	return models.ReservationCar{
		StartReservation: func() primitive.DateTime {
			t, _ := time.Parse("2006-01-02T15:04:05Z07:00", input.StartReservation) 
			return primitive.DateTime(t.UnixMilli())
		}(),
		EndReservation: func() primitive.DateTime {
			t, _ := time.Parse("2006-01-02T15:04:05Z07:00", input.EndReservation)
			return primitive.DateTime(t.UnixMilli())
		}(),
		NumberPlate: input.NumberPlate,
	}
}
