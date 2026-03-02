package domain

type VehicleType struct {
	ID         int     `gorm:"primaryKey;column:id"`
	TypeName   string  `gorm:"column:type_name"`
	HourlyRate float64 `gorm:"column:hourly_rate"`
}
