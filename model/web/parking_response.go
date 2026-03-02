package web

import "time"

type ParkingResponse struct {
	ID          int        `json:"id"`
	PlateNumber string     `json:"plate_number"`
	VehicleType string     `json:"vehicle_type"`
	CheckIn     time.Time  `json:"check_in"`
	CheckOut    *time.Time `json:"check_out,omitempty"`
	Duration    int        `json:"duration_hours,omitempty"`
	TotalFee    float64    `json:"total_fee"`
	PaidAt      *time.Time `json:"paid_at,omitempty"`
}
