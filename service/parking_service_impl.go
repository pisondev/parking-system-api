package service

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pisondev/parking-system-api/helper"
	"github.com/pisondev/parking-system-api/model/domain"
	"github.com/pisondev/parking-system-api/model/web"
	"github.com/pisondev/parking-system-api/repository"
	"gorm.io/gorm"
)

type ParkingServiceImpl struct {
	ParkingRepository repository.ParkingRepository
	DB                *gorm.DB
}

func NewParkingService(parkingRepository repository.ParkingRepository, db *gorm.DB) ParkingService {
	return &ParkingServiceImpl{
		ParkingRepository: parkingRepository,
		DB:                db,
	}
}

func (s *ParkingServiceImpl) Create(ctx context.Context, request web.ParkingCreateRequest) (web.ParkingResponse, error) {
	_, err := s.ParkingRepository.FindActiveByVehicleID(ctx, s.DB, request.VehicleID)
	if err == nil {
		return web.ParkingResponse{}, fiber.NewError(fiber.StatusBadRequest, "the vehicle is parked and has not been checked out")
	}

	tx := domain.ParkingTransaction{
		VehicleID: request.VehicleID,
		CheckIn:   time.Now(),
	}

	result, err := s.ParkingRepository.Save(ctx, s.DB, tx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return web.ParkingResponse{}, fiber.NewError(fiber.StatusBadRequest, "vehicle is already parked (unique constraint violation)")
		}
		return web.ParkingResponse{}, err
	}

	completeData, _ := s.ParkingRepository.FindById(ctx, s.DB, result.ID)
	return helper.ToParkingResponse(completeData), nil
}

func (s *ParkingServiceImpl) UpdateCheckout(ctx context.Context, id int) (web.ParkingResponse, error) {
	var response web.ParkingResponse

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		parking, err := s.ParkingRepository.FindById(ctx, tx, id)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "parking transaction not found")
		}

		if parking.CheckOut != nil {
			return fiber.NewError(fiber.StatusBadRequest, "this transaction has been completed previously")
		}

		checkOutTime := time.Now()
		duration := checkOutTime.Sub(parking.CheckIn)

		hours := math.Ceil(duration.Hours())
		if hours < 1 {
			hours = 1
		}
		totalFee := hours * parking.Vehicle.VehicleType.HourlyRate

		rowsAffected, err := s.ParkingRepository.Checkout(ctx, tx, id, checkOutTime, totalFee, checkOutTime)
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return fiber.NewError(fiber.StatusBadRequest, "checkout failed: race condition detected or invalid transaction")
		}

		parking.CheckOut = &checkOutTime
		parking.TotalFee = totalFee
		parking.PaidAt = &checkOutTime
		response = helper.ToParkingResponse(parking)

		return nil
	})

	return response, err
}

func (s *ParkingServiceImpl) FindById(ctx context.Context, id int) (web.ParkingResponse, error) {
	parking, err := s.ParkingRepository.FindById(ctx, s.DB, id)
	if err != nil {
		return web.ParkingResponse{}, fiber.NewError(fiber.StatusNotFound, "parking transaction not found")
	}
	return helper.ToParkingResponse(parking), nil
}

func (s *ParkingServiceImpl) FindAll(ctx context.Context) ([]web.ParkingResponse, error) {
	parkings, err := s.ParkingRepository.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}
	return helper.ToParkingResponses(parkings), nil
}

func (s *ParkingServiceImpl) Delete(ctx context.Context, id int) error {
	parking, err := s.ParkingRepository.FindById(ctx, s.DB, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "parking transaction not found")
	}

	if parking.CheckOut == nil {
		return fiber.NewError(fiber.StatusBadRequest, "don't delete history: vehicle is still in the parking area")
	}

	return s.ParkingRepository.Delete(ctx, s.DB, parking)
}
