package repository

import (
	"context"
	"time"

	"github.com/pisondev/parking-system-api/model/domain"
	"gorm.io/gorm"
)

type ParkingRepository interface {
	Save(ctx context.Context, db *gorm.DB, parking domain.ParkingTransaction) (domain.ParkingTransaction, error)
	Delete(ctx context.Context, db *gorm.DB, parking domain.ParkingTransaction) error
	FindById(ctx context.Context, db *gorm.DB, id int) (domain.ParkingTransaction, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.ParkingTransaction, error)
	FindActiveByVehicleID(ctx context.Context, db *gorm.DB, vehicleID int) (domain.ParkingTransaction, error)
	Checkout(ctx context.Context, db *gorm.DB, id int, checkOut time.Time, totalFee float64, paidAt time.Time) (int64, error)
}
