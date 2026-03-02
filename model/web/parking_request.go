package web

type ParkingCreateRequest struct {
	VehicleID int `json:"vehicle_id" validate:"required"`
}
