package repository

import (
	"context"
	"time"

	"github.com/pisondev/parking-system-api/model/domain"
	"gorm.io/gorm"
)

type ParkingRepositoryImpl struct {
}

func NewParkingRepository() ParkingRepository {
	return &ParkingRepositoryImpl{}
}

func (r *ParkingRepositoryImpl) Save(ctx context.Context, db *gorm.DB, parking domain.ParkingTransaction) (domain.ParkingTransaction, error) {
	err := db.WithContext(ctx).Create(&parking).Error
	return parking, err
}

func (r *ParkingRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, parking domain.ParkingTransaction) error {
	return db.WithContext(ctx).Delete(&parking).Error
}

func (r *ParkingRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, id int) (domain.ParkingTransaction, error) {
	var parking domain.ParkingTransaction
	err := db.WithContext(ctx).
		Preload("Vehicle").
		Preload("Vehicle.VehicleType").
		First(&parking, id).Error
	return parking, err
}

func (r *ParkingRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.ParkingTransaction, error) {
	var parkings []domain.ParkingTransaction
	err := db.WithContext(ctx).
		Preload("Vehicle").
		Preload("Vehicle.VehicleType").
		Find(&parkings).Error
	return parkings, err
}

func (r *ParkingRepositoryImpl) FindActiveByVehicleID(ctx context.Context, db *gorm.DB, vehicleID int) (domain.ParkingTransaction, error) {
	var parking domain.ParkingTransaction
	err := db.WithContext(ctx).
		Where("vehicle_id = ? AND check_out IS NULL", vehicleID).
		First(&parking).Error
	return parking, err
}

func (r *ParkingRepositoryImpl) Checkout(ctx context.Context, db *gorm.DB, id int, checkOut time.Time, totalFee float64, paidAt time.Time) (int64, error) {
	result := db.WithContext(ctx).
		Model(&domain.ParkingTransaction{}).
		Where("id = ? AND check_out IS NULL", id).
		Updates(map[string]interface{}{
			"check_out": checkOut,
			"total_fee": totalFee,
			"paid_at":   paidAt,
		})

	return result.RowsAffected, result.Error
}
