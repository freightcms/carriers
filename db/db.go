package db

import (
	"context"

	"github.com/freightcms/carriers/models"
)

type CarrierDb interface {
	GetCarriers(ctx context.Context) ([]models.FreightCarrierModel, error)
	GetCarrier(ctx context.Context, id string) (*models.FreightCarrierModel, error)
	CreateCarrier(ctx context.Context, carrier *models.FreightCarrierModel) error
	UpdateCarrier(ctx context.Context, id string, carrier *models.FreightCarrierModel) error
	DeleteCarrier(ctx context.Context, id string) error
}
