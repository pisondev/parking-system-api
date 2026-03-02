package service

import (
	"context"

	"github.com/pisondev/parking-system-api/model/web"
)

type ParkingService interface {
	Create(ctx context.Context, request web.ParkingCreateRequest) (web.ParkingResponse, error)
	UpdateCheckout(ctx context.Context, id int) (web.ParkingResponse, error)
	FindById(ctx context.Context, id int) (web.ParkingResponse, error)
	FindAll(ctx context.Context) ([]web.ParkingResponse, error)
	Delete(ctx context.Context, id int) error
}
