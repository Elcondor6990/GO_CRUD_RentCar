package dependencies

import (
	"Go_Rent_Car/src/config"
	"Go_Rent_Car/src/repository"
	"Go_Rent_Car/src/service"
	"os"
)

type AppDependencies struct {
	ReservationService service.ReservationCarService
}

func Init() *AppDependencies {

	db := config.MongoConnection(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB_NAME"),
	)

	reserveRepo := repository.NewReservationCarRepository(db)
	reservationService := service.NewReservationCarService(reserveRepo)

	return &AppDependencies{
		ReservationService: reservationService,
	}
}