package service

import (
	"Go_Rent_Car/src/dto"
	mapper "Go_Rent_Car/src/factory"
	"Go_Rent_Car/src/repository"
	"context"

	"github.com/gofiber/fiber/v3/log"
)

type reservationCarServiceImpl struct {
	repo repository.ReservationCarRepository
}

type ReservationCarService interface{
	Create(ctx context.Context, input dto.ReservationCarDto)(*dto.ReservationCarDto, error)
	GetByID(ctx context.Context, id string)(*dto.ReservationCarDto, error)
	Delete(ctx context.Context, id string) error
}

func NewReservationCarService(repo repository.ReservationCarRepository) ReservationCarService{
	return &reservationCarServiceImpl{repo: repo}
}


func (s *reservationCarServiceImpl) Create(ctx context.Context, input dto.ReservationCarDto)(*dto.ReservationCarDto, error){
	reservation := mapper.DtoToModel(input)

	createdReserve, err := s.repo.Create(ctx, &reservation)
	if err != nil{
		log.Errorf("service: error creating reservation: %v", err)
		return nil, err
	}

	if createdReserve.StartReservation > createdReserve.EndReservation{
		log.Errorf("service: error: the beginning must be greater than the end, detail: %v", err)
		return nil, err
	}

	log.Info("service: reservation created successfully")
	reservationDto := mapper.ModelToDto(*createdReserve)

	return &reservationDto, nil
}


func (s *reservationCarServiceImpl) GetByID(ctx context.Context, id string)(*dto.ReservationCarDto, error){
	reservation, err := s.repo.FindByID(ctx, id)
	if err != nil{
		log.Errorf("service: error getting reservation by ID: %v", err)
		return nil, err
	}

	log.Infof("service: reservation found with ID: %s", reservation.ID)
	foundedReserve := mapper.ModelToDto(*reservation)
	return &foundedReserve, nil
}

func (s *reservationCarServiceImpl) Delete(ctx context.Context, id string) error{
	err := s.repo.Delete(ctx, id)
	if err != nil{
		log.Errorf("service: error deleting reservation: %v", err)
		return err
	}
	
	log.Info("service: reservation deleted successfully")
	return nil
}