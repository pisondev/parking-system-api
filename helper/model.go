package helper

import (
	"math"

	"github.com/pisondev/parking-system-api/model/domain"
	"github.com/pisondev/parking-system-api/model/web"
)

func ToParkingResponse(tx domain.ParkingTransaction) web.ParkingResponse {
	response := web.ParkingResponse{
		ID:          tx.ID,
		PlateNumber: tx.Vehicle.PlateNumber,
		VehicleType: tx.Vehicle.VehicleType.TypeName,
		CheckIn:     tx.CheckIn,
		CheckOut:    tx.CheckOut,
		TotalFee:    tx.TotalFee,
		PaidAt:      tx.PaidAt,
	}

	if tx.CheckOut != nil {
		duration := tx.CheckOut.Sub(tx.CheckIn)
		hours := int(math.Ceil(duration.Hours()))
		if hours < 1 {
			hours = 1
		}
		response.Duration = hours
	}

	return response
}

func ToParkingResponses(transactions []domain.ParkingTransaction) []web.ParkingResponse {
	var parkingResponses []web.ParkingResponse
	for _, tx := range transactions {
		parkingResponses = append(parkingResponses, ToParkingResponse(tx))
	}
	return parkingResponses
}
