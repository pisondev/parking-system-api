package domain

import "time"

type ParkingTransaction struct {
	ID        int        `gorm:"primaryKey;column:id"`
	VehicleID int        `gorm:"column:vehicle_id"`
	Vehicle   Vehicle    `gorm:"foreignKey:VehicleID;references:ID"`
	CheckIn   time.Time  `gorm:"column:check_in;autoCreateTime"`
	CheckOut  *time.Time `gorm:"column:check_out"`
	TotalFee  float64    `gorm:"column:total_fee"`
	PaidAt    *time.Time `gorm:"column:paid_at"`
}
