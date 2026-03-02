package domain

type Vehicle struct {
	ID            int         `gorm:"primaryKey;column:id"`
	PlateNumber   string      `gorm:"column:plate_number"`
	VehicleTypeID int         `gorm:"column:vehicle_type_id"`
	VehicleType   VehicleType `gorm:"foreignKey:VehicleTypeID;references:ID"`
}
